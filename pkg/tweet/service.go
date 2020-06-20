package tweet

import (
	"com.capturetweet/internal/convert"
	"com.capturetweet/pkg/service"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"net/url"
)

type serviceImpl struct {
	repo       Repository
	twitterAPI *anaconda.TwitterApi
	search     service.SearchService
	user       service.UserService
}

func NewService(repo Repository, search service.SearchService, userService service.UserService, twitterAPI *anaconda.TwitterApi) service.TweetService {
	return &serviceImpl{repo, twitterAPI, search, userService}
}

func (s serviceImpl) FindById(id string) (*service.TweetModel, error) {
	tweet, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	var resources []service.ResourceModel

	for _, res := range tweet.Resources {
		resources = append(resources, service.ResourceModel{
			ID:           res.ID,
			URL:          res.URL,
			Width:        res.Width,
			Height:       res.Height,
			ResourceType: res.URL,
		})
	}

	return &service.TweetModel{
		ID:              tweet.ID,
		PostedAt:        convert.Time(tweet.PostedAt),
		FullText:        tweet.FullText,
		Lang:            tweet.Lang,
		CaptureURL:      tweet.CaptureURL,
		CaptureThumbURL: tweet.CaptureThumbURL,
		FavoriteCount:   tweet.FavoriteCount,
		RetweetCount:    tweet.RetweetCount,
		AuthorID:        tweet.AuthorID,
		Resources:       resources,
	}, nil
}

func (s serviceImpl) Store(tweetURL string) (string, error) {
	tweetID, _, err := parseTweetURL(tweetURL)
	if err != nil {
		return "", err
	}

	tweetIdStr := fmt.Sprintf("%d", tweetID)
	if s.repo.Exist(tweetIdStr) {
		return tweetIdStr, nil
	}

	tweet, err := s.twitterAPI.GetTweet(tweetID, url.Values{})
	if err != nil {
		return "", err
	}

	err = s.repo.Store(&tweet)
	if err != nil {
		return "", err
	}

	go func(author *anaconda.User) {
		s.user.FindOrCreate(author)
	}(&tweet.User)

	go func(t anaconda.Tweet) {
		s.search.Put(t.IdStr, t.FullText, t.User.ScreenName)
	}(tweet)

	return tweetIdStr, nil
}

func (s serviceImpl) Search(term string, size, start, page int) ([]service.TweetModel, error) {
	searchModels, err := s.search.Search(term, size)
	if err != nil {
		return nil, err
	}

	var ids []string
	for _, model := range searchModels {
		ids = append(ids, model.TweetID)
	}

	tweets, err := s.repo.FindByIds(ids)
	if err != nil {
		return nil, err
	}

	var res []service.TweetModel

	for _, tweet := range tweets {
		var resources []service.ResourceModel

		for _, res := range tweet.Resources {
			resources = append(resources, service.ResourceModel{
				ID:           res.ID,
				URL:          res.URL,
				Width:        res.Width,
				Height:       res.Height,
				ResourceType: res.URL,
			})
		}

		res = append(res, service.TweetModel{
			ID:              tweet.ID,
			FullText:        tweet.FullText,
			Lang:            tweet.Lang,
			CaptureURL:      nil,
			CaptureThumbURL: nil,
			FavoriteCount:   tweet.FavoriteCount,
			RetweetCount:    tweet.RetweetCount,
			AuthorID:        tweet.AuthorID,
			PostedAt:        convert.Time(tweet.PostedAt),
			Resources:       resources,
		})
	}

	return res, nil
}

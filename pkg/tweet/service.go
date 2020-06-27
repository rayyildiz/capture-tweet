package tweet

import (
	"com.capturetweet/internal/convert"
	"com.capturetweet/pkg/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"go.uber.org/zap"
	"gocloud.dev/pubsub"
	"math/rand"
	"net/url"
	"time"
)

type serviceImpl struct {
	repo       Repository
	twitterAPI *anaconda.TwitterApi
	search     service.SearchService
	user       service.UserService
	log        *zap.Logger
	topic      *pubsub.Topic
}

func NewService(repo Repository, search service.SearchService, userService service.UserService, twitterAPI *anaconda.TwitterApi, log *zap.Logger, topic *pubsub.Topic) service.TweetService {
	return &serviceImpl{repo, twitterAPI, search, userService, log, topic}
}

func (s serviceImpl) FindById(id string) (*service.TweetModel, error) {
	tweet, err := s.repo.FindById(id)
	if err != nil {
		s.log.Error("tweet:service findById", zap.String("tweet_id", id), zap.Error(err))
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
	tweetID, tweetAuthor, err := parseTweetURL(tweetURL)
	if err != nil {
		s.log.Error("tweet:service store, parseTweetUrl", zap.String("url", tweetURL), zap.Error(err))
		return "", err
	}

	tweetIdStr := fmt.Sprintf("%d", tweetID)
	if s.repo.Exist(tweetIdStr) {
		return tweetIdStr, nil
	}

	tweet, err := s.twitterAPI.GetTweet(tweetID, url.Values{})
	if err != nil {
		s.log.Error("tweet:service store, getTweet", zap.Int64("tweet_id", tweetID), zap.Error(err))
		return "", err
	}

	err = s.repo.Store(&tweet)
	if err != nil {
		s.log.Error("tweet:service store, repo.store", zap.Int64("tweet_id", tweetID), zap.Error(err))
		return "", err
	}

	if _, err := s.user.FindOrCreate(&tweet.User); err != nil {
		s.log.Warn("tweet:service store, user.findOrCreate", zap.String("tweet_user", tweet.User.ScreenName), zap.Error(err))
	}

	go func(t anaconda.Tweet) {
		if err := s.search.Put(t.IdStr, t.FullText, t.User.ScreenName); err != nil {
			s.log.Warn("tweet:service store, search.put", zap.String("tweet_id", t.IdStr), zap.String("tweet_user", t.User.ScreenName), zap.Error(err))
		}
	}(tweet)

	go func(id, author, url string) {

		data, err := json.Marshal(service.CaptureRequestModel{
			ID:     id,
			Author: author,
			Url:    url,
		})
		if err != nil {
			s.log.Warn("tweet:service store, send pubsub message", zap.String("tweet_id", id), zap.String("tweet_user", author), zap.String("url", url), zap.Error(err))
			return
		}

		err = s.topic.Send(context.Background(), &pubsub.Message{
			Metadata: map[string]string{
				"tweet_id":   id,
				"tweet_user": author,
				"version":    "beta",
			},
			Body: data,
		})
		if err != nil {
			s.log.Warn("tweet:service store, send pubsub message", zap.String("tweet_id", id), zap.String("tweet_user", author), zap.String("url", url), zap.Error(err))
		} else {
			s.log.Info("tweet:service store, sent pubsub message", zap.String("tweet_id", id), zap.String("tweet_user", author), zap.String("url", url))
		}
	}(tweetIdStr, tweetAuthor, tweetURL)

	return tweetIdStr, nil
}

func (s serviceImpl) Search(term string, size, start, page int) ([]service.TweetModel, error) {
	searchModels, err := s.search.Search(term, size)
	if err != nil {
		s.log.Error("tweet:service search, search service call", zap.String("search_term", term), zap.Error(err))
		return nil, err
	}

	var ids []string
	for _, model := range searchModels {
		ids = append(ids, model.TweetID)
	}

	tweets, err := s.repo.FindByIds(ids)
	if err != nil {
		s.log.Error("tweet:service search, findByIds", zap.Strings("tweet_ids", ids), zap.Error(err))
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

func (s serviceImpl) UpdateCaptureImage(id, captureUrl, captureThumbUrl string) error {
	err := s.repo.UpdateCaptureURLs(id, captureUrl, captureThumbUrl)
	if err != nil {
		s.log.Error("tweet:service updateCaptureImage, findById", zap.String("tweet_id", id), zap.Error(err))
		return err
	}
	s.log.Info("tweet:service updateCaptureImage, updated capture images", zap.String("tweet_id", id))
	return nil
}

func (s serviceImpl) WatchChange(ctx context.Context, id string) (<-chan *service.TweetModel, error) {
	// TODO implement
	ch := make(chan *service.TweetModel)

	go func() {
		for {
			rand.Seed(time.Now().Unix())

			ch <- &service.TweetModel{
				ID:            "1272545509736017920",
				FullText:      fmt.Sprintf("test : %v", time.Now()),
				Lang:          "en",
				FavoriteCount: rand.Intn(100),
				RetweetCount:  rand.Intn(200),
			}
			time.Sleep(time.Second * 10)
		}
	}()

	return ch, nil
}

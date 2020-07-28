package tweet

import (
	"com.capturetweet/api"
	"com.capturetweet/internal/convert"
	"com.capturetweet/internal/infra"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"go.uber.org/zap"
	"gocloud.dev/pubsub"
	"net/url"
)

type serviceImpl struct {
	repo       Repository
	twitterAPI infra.TweetAPI
	search     api.SearchService
	user       api.UserService
	log        *zap.Logger
	topic      *pubsub.Topic
}

func NewService(repo Repository, search api.SearchService, userService api.UserService, twitterAPI infra.TweetAPI, log *zap.Logger, topic *pubsub.Topic) api.TweetService {
	return &serviceImpl{repo, twitterAPI, search, userService, log, topic}
}

func (s serviceImpl) FindById(ctx context.Context, id string) (*api.TweetModel, error) {
	tweet, err := s.repo.FindById(ctx, id)
	if err != nil {
		s.log.Error("tweet:service findById", zap.String("tweet_id", id), zap.Error(err))
		return nil, err
	}

	var resources []api.ResourceModel

	for _, res := range tweet.Resources {
		resources = append(resources, api.ResourceModel{
			ID:           res.ID,
			URL:          res.URL,
			Width:        res.Width,
			Height:       res.Height,
			ResourceType: res.URL,
		})
	}

	return &api.TweetModel{
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

func (s serviceImpl) Store(ctx context.Context, tweetURL string) (string, error) {
	tweetID, tweetAuthor, err := parseTweetURL(tweetURL)
	if err != nil {
		s.log.Error("tweet:service store, parseTweetUrl", zap.String("url", tweetURL), zap.Error(err))
		return "", err
	}

	tweetIdStr := fmt.Sprintf("%d", tweetID)
	if s.repo.Exist(ctx, tweetIdStr) {
		return tweetIdStr, nil
	}

	tweet, err := s.twitterAPI.GetTweet(tweetID, url.Values{})
	if err != nil {
		s.log.Error("tweet:service store, getTweet", zap.Int64("tweet_id", tweetID), zap.Error(err))
		return "", err
	}

	err = s.repo.Store(ctx, &tweet)
	if err != nil {
		s.log.Error("tweet:service store, repo.store", zap.Int64("tweet_id", tweetID), zap.Error(err))
		return "", err
	}

	if _, err := s.user.FindOrCreate(ctx, &tweet.User); err != nil {
		s.log.Warn("tweet:service store, user.findOrCreate", zap.String("tweet_user", tweet.User.ScreenName), zap.Error(err))
	}

	go func(t anaconda.Tweet) {
		if err := s.search.Put(ctx, t.IdStr, t.FullText, t.User.ScreenName); err != nil {
			s.log.Warn("tweet:service store, search.put", zap.String("tweet_id", t.IdStr), zap.String("tweet_user", t.User.ScreenName), zap.Error(err))
		}
	}(tweet)

	go func(id, author, url string) {

		data, err := json.Marshal(api.CaptureRequestModel{
			ID:     id,
			Author: author,
			Url:    url,
		})
		if err != nil {
			s.log.Warn("tweet:service store, send pubsub message", zap.String("tweet_id", id), zap.String("tweet_user", author), zap.String("url", url), zap.Error(err))
			return
		}

		err = s.topic.Send(ctx, &pubsub.Message{
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

func (s serviceImpl) Search(ctx context.Context, term string, size, start, page int) ([]api.TweetModel, error) {
	searchModels, err := s.search.Search(ctx, term, size)
	if err != nil {
		s.log.Error("tweet:service search, search service call", zap.String("search_term", term), zap.Error(err))
		return nil, err
	}

	var ids []string
	for _, model := range searchModels {
		ids = append(ids, model.TweetID)
	}

	tweets, err := s.repo.FindByIds(ctx, ids)
	if err != nil {
		s.log.Error("tweet:service search, findByIds", zap.Strings("tweet_ids", ids), zap.Error(err))
		return nil, err
	}

	var list []api.TweetModel

	for _, tweet := range tweets {
		list = append(list, convertToTweet(tweet))
	}

	return list, nil
}

func (s serviceImpl) UpdateLargeImage(ctx context.Context, id, captureUrl string) error {
	err := s.repo.UpdateLargeImage(ctx, id, captureUrl)
	if err != nil {
		s.log.Error("tweet:service updateLargeImage, findById", zap.String("tweet_id", id), zap.Error(err))
		return err
	}
	s.log.Info("tweet:service updateLargeImage, updated capture images", zap.String("tweet_id", id))
	return nil
}

func (s serviceImpl) UpdateThumbImage(ctx context.Context, id, captureUrl string) error {
	err := s.repo.UpdateThumbImage(ctx, id, captureUrl)
	if err != nil {
		s.log.Error("tweet:service updateThumbImage, findById", zap.String("tweet_id", id), zap.Error(err))
		return err
	}
	s.log.Info("tweet:service updateThumbImage, updated capture images", zap.String("tweet_id", id))
	return nil
}

func (s serviceImpl) SearchByUser(ctx context.Context, userId string) ([]api.TweetModel, error) {
	tweets, err := s.repo.FindByUser(ctx, userId)
	if err != nil {
		s.log.Error("tweet:service searchByUser, findById", zap.String("user_id", userId), zap.Error(err))
		return nil, err
	}

	var list []api.TweetModel

	for _, tweet := range tweets {
		list = append(list, convertToTweet(tweet))
	}

	return list, nil
}

func convertToTweet(tweet Tweet) api.TweetModel {
	var resources []api.ResourceModel

	for _, res := range tweet.Resources {
		resources = append(resources, api.ResourceModel{
			ID:           res.ID,
			URL:          res.URL,
			Width:        res.Width,
			Height:       res.Height,
			ResourceType: res.URL,
		})
	}

	return api.TweetModel{
		ID:              tweet.ID,
		FullText:        tweet.FullText,
		Lang:            tweet.Lang,
		CaptureURL:      tweet.CaptureURL,
		CaptureThumbURL: tweet.CaptureThumbURL,
		FavoriteCount:   tweet.FavoriteCount,
		RetweetCount:    tweet.RetweetCount,
		AuthorID:        tweet.AuthorID,
		PostedAt:        convert.Time(tweet.PostedAt),
		Resources:       resources,
	}
}

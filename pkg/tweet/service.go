package tweet

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/url"

	"capturetweet.com/api"
	"capturetweet.com/internal/convert"
	"capturetweet.com/internal/infra"
	"github.com/ChimeraCoder/anaconda"
	"gocloud.dev/pubsub"
)

type serviceImpl struct {
	repo       Repository
	twitterAPI infra.TweetAPI
	search     api.SearchService
	user       api.UserService
	topic      *pubsub.Topic
}

func NewService(repo Repository, search api.SearchService, userService api.UserService, twitterAPI infra.TweetAPI, topic *pubsub.Topic) api.TweetService {
	return &serviceImpl{
		repo:       repo,
		twitterAPI: twitterAPI,
		search:     search,
		user:       userService,
		topic:      topic,
	}
}

func NewServiceWithRepository(repo Repository) api.TweetService {
	return NewService(repo, nil, nil, nil, nil)
}

func (s serviceImpl) FindById(ctx context.Context, id string) (*api.TweetModel, error) {
	tweet, err := s.repo.FindById(ctx, id)
	if err != nil {
		slog.Error("tweet:service findById", slog.String("tweet_id", id), slog.Any("err", err))
		return nil, err
	}

	var resources []api.ResourceModel

	for _, res := range tweet.Resources {
		resources = append(resources, api.ResourceModel{
			ID:           res.ID,
			URL:          res.URL,
			Width:        res.Width,
			Height:       res.Height,
			ResourceType: res.MediaType,
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
		slog.Error("tweet:service store, parseTweetUrl", slog.String("url", tweetURL), slog.Any("err", err))
		return "", err
	}

	tweetIdStr := fmt.Sprintf("%d", tweetID)
	if s.repo.Exist(ctx, tweetIdStr) {
		return tweetIdStr, nil
	}
	tweet, err := s.twitterAPI.GetTweet(tweetID, url.Values{})

	if err != nil {
		slog.Error("tweet:service store, getTweet", slog.Int64("tweet_id", tweetID), slog.Any("err", err))
		return "", err
	}

	err = s.repo.Store(ctx, &tweet)
	if err != nil {
		slog.Error("tweet:service store, repo.store", slog.Int64("tweet_id", tweetID), slog.Any("err", err))
		return "", err
	}

	if _, err := s.user.FindOrCreate(ctx, &tweet.User); err != nil {
		slog.Warn("tweet:service store, user.findOrCreate", slog.String("tweet_user", tweet.User.ScreenName), slog.Any("err", err))
	}

	go func(ctx context.Context, t anaconda.Tweet) {
		if err := s.search.Put(ctx, t.IdStr, t.FullText, t.User.ScreenName); err != nil {
			slog.Warn("tweet:service store, search.put", slog.String("tweet_id", t.IdStr), slog.String("tweet_user", t.User.ScreenName), slog.Any("err", err))
		}
	}(ctx, tweet)

	go func(ctx context.Context, id, author, url string) {

		data, err := json.Marshal(api.CaptureRequestModel{
			ID:     id,
			Author: author,
			Url:    url,
		})
		if err != nil {
			slog.Warn("tweet:service store, send pubsub message", slog.String("tweet_id", id), slog.String("tweet_user", author), slog.String("url", url), slog.Any("err", err))
			return
		}
		err = s.topic.Send(ctx, &pubsub.Message{
			Metadata: map[string]string{
				"tweet_id":            id,
				"tweet_user":          author,
				"version":             "prod",
				"messaging.system":    "pubsub",
				"messaging.operation": "send",
			},
			Body: data,
		})

		if err != nil {
			slog.Warn("tweet:service store, send pubsub message", slog.String("tweet_id", id), slog.String("tweet_user", author), slog.String("url", url), slog.Any("err", err))
		} else {
			slog.Info("tweet:service store, sent pubsub message", slog.String("tweet_id", id), slog.String("tweet_user", author), slog.String("url", url))
		}
	}(ctx, tweetIdStr, tweetAuthor, tweetURL)

	return tweetIdStr, nil
}

func (s serviceImpl) Search(ctx context.Context, term string, size, start, page int) ([]api.TweetModel, error) {
	searchModels, err := s.search.Search(ctx, term, size)
	if err != nil {
		slog.Error("tweet:service search, search service call", slog.String("search_term", term), slog.Any("err", err))
		return nil, err
	}

	var ids []string
	for _, model := range searchModels {
		ids = append(ids, model.TweetID)
	}

	tweets, err := s.repo.FindByIds(ctx, ids)
	if err != nil {
		slog.Error("tweet:service search, findByIds", slog.Any("tweet_ids", ids), slog.Any("err", err))
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
		slog.Error("tweet:service updateLargeImage, findById", slog.String("tweet_id", id), slog.Any("err", err))
		return err
	}
	slog.Info("tweet:service updateLargeImage, updated capture images", slog.String("tweet_id", id))
	return nil
}

func (s serviceImpl) UpdateThumbImage(ctx context.Context, id, captureUrl string) error {
	err := s.repo.UpdateThumbImage(ctx, id, captureUrl)
	if err != nil {
		slog.Error("tweet:service updateThumbImage, findById", slog.String("tweet_id", id), slog.Any("err", err))
		return fmt.Errorf("update thumnmail,%w", err)
	}
	slog.Info("tweet:service updateThumbImage, updated capture images", slog.String("tweet_id", id))
	return nil
}

func (s serviceImpl) SearchByUser(ctx context.Context, userId string) ([]api.TweetModel, error) {
	tweets, err := s.repo.FindByUser(ctx, userId)
	if err != nil {
		slog.Error("tweet:service searchByUser, findById", slog.String("user_id", userId), slog.Any("err", err))
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

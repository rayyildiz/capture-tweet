package graph

import (
	"com.capturetweet/internal/convert"
	"com.capturetweet/pkg/service"
	"context"
	"go.uber.org/zap"
)

type subscriptionImpl struct {
}

func newSubscriptionResolver() SubscriptionResolver {
	return &subscriptionImpl{}
}

func (subscriptionImpl) Captured(ctx context.Context, id string) (<-chan *Tweet, error) {
	ch := make(chan *Tweet, 1)

	watchChange, err := _twitterService.WatchChange(ctx, id)

	if err != nil {
		_log.Error("subscription capture", zap.String("tweet_id", id), zap.Error(err))
		return nil, err
	}

	go func(modelCh <-chan *service.TweetModel) {

		for model := range modelCh {
			ch <- &Tweet{
				ID:              model.ID,
				FullText:        model.FullText,
				PostedAt:        model.PostedAt,
				AuthorID:        convert.String(model.AuthorID),
				CaptureURL:      model.CaptureURL,
				CaptureThumbURL: model.CaptureThumbURL,
				FavoriteCount:   convert.Int(model.FavoriteCount),
				Lang:            convert.String(model.Lang),
				RetweetCount:    convert.Int(model.RetweetCount),
			}
		}

	}(watchChange)

	return ch, nil
}

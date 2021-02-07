package resolver

import (
	"com.capturetweet/api"
	"com.capturetweet/internal/convert"
	"context"
	"errors"
	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
	"gocloud.dev/gcerrors"
)

type queryResolverImpl struct {
}

func newQueryResolver() QueryResolver {
	return &queryResolverImpl{}
}

func (r queryResolverImpl) Tweet(ctx context.Context, id string) (*Tweet, error) {
	ctx, span := _tracer.Start(ctx, "tweet")
	defer span.End()

	model, err := _twitterService.FindById(ctx, id)
	code := gcerrors.Code(err)
	if code == gcerrors.NotFound {
		zap.L().Warn("tweet not found", zap.String("id", id))
		return nil, nil
	}

	if err != nil {
		sentry.CaptureException(err)
		zap.L().Error("tweet find error", zap.String("id", id), zap.Error(err))
		return nil, err
	}
	var resources []*Resource
	for _, res := range model.Resources {
		resources = append(resources, &Resource{
			ID:        res.ID,
			URL:       res.URL,
			MediaType: convert.String(res.ResourceType),
			Width:     convert.Int(res.Width),
			Height:    convert.Int(res.Height),
		})
	}

	return &Tweet{
		ID:              model.ID,
		FullText:        model.FullText,
		PostedAt:        model.PostedAt,
		CaptureURL:      model.CaptureURL,
		CaptureThumbURL: model.CaptureThumbURL,
		FavoriteCount:   convert.Int(model.FavoriteCount),
		Lang:            convert.String(model.Lang),
		RetweetCount:    convert.Int(model.RetweetCount),
		AuthorID:        convert.String(model.AuthorID),
		Resources:       resources,
	}, nil
}

func (r queryResolverImpl) SearchByUser(ctx context.Context, userID string) ([]*Tweet, error) {
	ctx, span := _tracer.Start(ctx, "searchByUser")
	defer span.End()

	models, err := _twitterService.SearchByUser(ctx, userID)
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Error("search by user error", zap.String("user_id", userID), zap.Error(err))
		return nil, errors.New("could not ind any result")
	}
	var list []*Tweet
	for _, model := range models {
		list = append(list, convertToModel(model))
	}

	return list, nil
}

func (r queryResolverImpl) Search(ctx context.Context, input SearchInput, size int, page int, start int) ([]*Tweet, error) {
	ctx, span := _tracer.Start(ctx, "search")
	defer span.End()

	models, err := _twitterService.Search(ctx, input.Term, size, start, page)
	if err != nil {
		sentry.CaptureException(err)
		zap.L().Error("search error", zap.String("term", input.Term), zap.Error(err))
		return nil, errors.New("could not ind any result")
	}
	var list []*Tweet
	for _, model := range models {
		list = append(list, convertToModel(model))
	}

	return list, nil
}

func convertToModel(model api.TweetModel) *Tweet {
	var resources []*Resource
	for _, res := range model.Resources {
		resources = append(resources, &Resource{
			ID:        res.ID,
			URL:       res.URL,
			MediaType: convert.String(res.ResourceType),
			Width:     convert.Int(res.Width),
			Height:    convert.Int(res.Height),
		})
	}

	return &Tweet{
		ID:              model.ID,
		FullText:        model.FullText,
		PostedAt:        model.PostedAt,
		CaptureURL:      model.CaptureURL,
		CaptureThumbURL: model.CaptureThumbURL,
		FavoriteCount:   convert.Int(model.FavoriteCount),
		Lang:            convert.String(model.Lang),
		RetweetCount:    convert.Int(model.RetweetCount),
		AuthorID:        convert.String(model.AuthorID),
		Resources:       resources,
	}
}

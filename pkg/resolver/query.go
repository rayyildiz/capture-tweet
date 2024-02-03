package resolver

import (
	"context"
	"errors"
	"log/slog"

	"capturetweet.com/api"
	"capturetweet.com/internal/convert"
	"gocloud.dev/gcerrors"
)

type queryResolverImpl struct {
}

func newQueryResolver() QueryResolver {
	return &queryResolverImpl{}
}

func (r queryResolverImpl) Tweet(ctx context.Context, id string) (*Tweet, error) {
	model, err := _twitterService.FindById(ctx, id)
	code := gcerrors.Code(err)
	if code == gcerrors.NotFound {
		slog.Warn("tweet not found", slog.String("id", id))
		return nil, nil
	}

	if err != nil {
		slog.Error("tweet find error", slog.String("id", id), slog.Any("err", err))
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
	models, err := _twitterService.SearchByUser(ctx, userID)
	if err != nil {
		slog.Error("search by user error", slog.String("user_id", userID), slog.Any("err", err))
		return nil, errors.New("could not ind any result")
	}
	var list []*Tweet
	for _, model := range models {
		list = append(list, convertToModel(model))
	}

	return list, nil
}

func (r queryResolverImpl) Search(ctx context.Context, input SearchInput, size int, page int, start int) ([]*Tweet, error) {
	models, err := _twitterService.Search(ctx, input.Term, size, start, page)
	if err != nil {
		slog.Error("search error", slog.String("term", input.Term), slog.Any("err", err))
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

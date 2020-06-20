package graph

import (
	"com.capturetweet/internal/convert"
	"context"
	"errors"
	"go.uber.org/zap"
	"gocloud.dev/gcerrors"
)

type queryResolverImpl struct {
}

func newQueryResolver() QueryResolver {
	return &queryResolverImpl{}
}

func (r queryResolverImpl) Tweet(ctx context.Context, id string) (*Tweet, error) {
	model, err := _twitterService.FindById(id)
	code := gcerrors.Code(err)
	if code == gcerrors.NotFound {
		_log.Warn("tweet not found", zap.String("id", id))
		return nil, nil
	}

	if err != nil {
		_log.Error("tweet find error", zap.String("id", id), zap.Error(err))
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

func (r queryResolverImpl) Search(ctx context.Context, input SearchInput, size int, page int, start int) ([]*Tweet, error) {

	models, err := _twitterService.Search(input.Term, size, start, page)
	if err != nil {
		_log.Error("search error", zap.String("term", input.Term), zap.Error(err))
		return nil, errors.New("could not ind any result")
	}
	var list []*Tweet
	for _, model := range models {

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

		list = append(list, &Tweet{
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
		})
	}

	return list, nil
}

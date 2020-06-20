package graph

import (
	"com.capturetweet/internal/convert"
	"context"
	"errors"
)

type queryResolverImpl struct {
}

func newQueryResolver() QueryResolver {
	return &queryResolverImpl{}
}

func (r queryResolverImpl) Tweet(ctx context.Context, id string) (*Tweet, error) {

	return nil, nil
}

func (r queryResolverImpl) Search(ctx context.Context, input SearchInput, size int, page int, start int) ([]*Tweet, error) {

	models, err := _twitterService.Search(input.Term, size, start, page)
	if err != nil {
		return nil, errors.New("could not ind any result")
	}
	var list []*Tweet
	for _, model := range models {
		list = append(list, &Tweet{
			ID:              model.ID,
			FullText:        model.FullText,
			PostedAt:        model.PostedAt,
			CaptureURL:      model.CaptureURL,
			CaptureThumbURL: model.CaptureThumbURL,
			FavoriteCount:   convert.Int(model.FavoriteCount),
			Lang:            convert.String(model.Lang),
			RetweetCount:    convert.Int(model.RetweetCount),
		})
	}

	return list, nil
}

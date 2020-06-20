package graph

import (
	"com.capturetweet/internal/convert"
	"context"
)

type mutationResolverImpl struct {
}

func newMutationResolver() MutationResolver {
	return &mutationResolverImpl{}
}

func (r mutationResolverImpl) Capture(ctx context.Context, url string) (*Tweet, error) {
	id, err := _twitterService.Store(url)
	if err != nil {
		return nil, err
	}

	model, err := _twitterService.FindById(id)
	if err != nil {
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
		AuthorID:        convert.String(model.AuthorID),
		CaptureURL:      model.CaptureURL,
		CaptureThumbURL: model.CaptureThumbURL,
		FavoriteCount:   convert.Int(model.FavoriteCount),
		Lang:            convert.String(model.Lang),
		RetweetCount:    convert.Int(model.RetweetCount),
		Resources:       resources,
	}, nil
}

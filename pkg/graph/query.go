package graph

import (
	"com.capturetweet/internal/convert"
	"com.capturetweet/pkg/graph/generated"
	"com.capturetweet/pkg/service"
	"context"
	"errors"
)

type queryResolverImpl struct {
	tweetSvc service.TweetService
}

func newQueryResolver(ts service.TweetService) generated.QueryResolver {
	return &queryResolverImpl{ts}
}

func (r queryResolverImpl) Tweet(ctx context.Context, id string) (*generated.Tweet, error) {

	return nil, nil
}

func (r queryResolverImpl) Search(ctx context.Context, input generated.SearchInput, size int, page int, start int) ([]*generated.Tweet, error) {

	models, err := r.tweetSvc.Search(input.Term, size, start, page)
	if err != nil {
		return nil, errors.New("could not ind any result")
	}
	var list []*generated.Tweet
	for _, model := range models {
		list = append(list, &generated.Tweet{
			ID:                 model.ID,
			FullText:           model.FullText,
			CreatedAt:          nil,
			User:               nil,
			ImageURL:           model.CaptureURL,
			ThumbnaildImageURL: model.ThumbnailURL,
			FavoriteCount:      convert.Int(model.FavoriteCount),
			RetweetCount:       convert.Int(model.RetweetCount),
			Resources:          nil,
		})
	}

	return list, nil
}

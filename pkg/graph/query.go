package graph

import (
	"com.capturetweet/pkg/graph/generated"
	"com.capturetweet/pkg/service"
	"context"
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

func (r queryResolverImpl) Search(ctx context.Context, input generated.SearchInput, size *int, cursor *string) ([]*generated.Tweet, error) {
	return nil, nil
}

package graph

import (
	"com.capturetweet/pkg/graph/generated"
	"com.capturetweet/pkg/service"
	"context"
)

type mutationResolverImpl struct {
	tweetSvc service.TweetService
}

func newMutationResolver(ts service.TweetService) generated.MutationResolver {
	return &mutationResolverImpl{ts}
}

func (r mutationResolverImpl) Capture(ctx context.Context, url *string) (*generated.Tweet, error) {

	return nil, nil
}

package graph

import (
	"context"
)

type mutationResolverImpl struct {
}

func newMutationResolver() MutationResolver {
	return &mutationResolverImpl{}
}

func (r mutationResolverImpl) Capture(ctx context.Context, url *string) (*Tweet, error) {

	return nil, nil
}

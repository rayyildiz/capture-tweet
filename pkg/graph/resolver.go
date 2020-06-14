//go:generate go run github.com/99designs/gqlgen
package graph

import (
	"com.capturetweet/pkg/graph/generated"
	"com.capturetweet/pkg/service"
)

type resolver struct {
	userSvc  service.UserService
	tweetSvc service.TweetService
}

func NewResolver(us service.UserService, ts service.TweetService) generated.ResolverRoot {
	return &resolver{us, ts}
}

func (r resolver) Mutation() generated.MutationResolver {
	return newMutationResolver(r.tweetSvc)
}

func (r resolver) Query() generated.QueryResolver {
	return newQueryResolver(r.tweetSvc)
}

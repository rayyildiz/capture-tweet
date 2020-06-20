//go:generate go run github.com/99designs/gqlgen
package graph

import (
	"com.capturetweet/pkg/service"
)

type resolver struct {
}

func NewResolver() ResolverRoot {
	return &resolver{}
}

func (r resolver) Mutation() MutationResolver {
	return newMutationResolver()
}

func (r resolver) Query() QueryResolver {
	return newQueryResolver()
}

var (
	_twitterService service.TweetService  = nil
	_searchService  service.SearchService = nil
	_userService    service.UserService   = nil
)

func InitService(twitterService service.TweetService, searchService service.SearchService, userService service.UserService) {
	_twitterService = twitterService
	_userService = userService
	_searchService = searchService
}

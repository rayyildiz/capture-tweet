//go:generate go run github.com/99designs/gqlgen
package resolver

import (
	"com.capturetweet/api"
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

func (r resolver) __Directive() __DirectiveResolver {
	return nil
}

var (
	_twitterService api.TweetService   = nil
	_userService    api.UserService    = nil
	_contentService api.ContentService = nil
)

func InitService(twitterService api.TweetService, userService api.UserService, contentService api.ContentService) {
	_twitterService = twitterService
	_userService = userService
	_contentService = contentService
}

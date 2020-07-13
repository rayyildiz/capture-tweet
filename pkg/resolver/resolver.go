//go:generate go run github.com/99designs/gqlgen
package resolver

import (
	"com.capturetweet/api"
	"go.uber.org/zap"
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
	_twitterService api.TweetService   = nil
	_userService    api.UserService    = nil
	_log            *zap.Logger        = nil
	_contentService api.ContentService = nil
)

func InitService(log *zap.Logger, twitterService api.TweetService, userService api.UserService, contentService api.ContentService) {
	_twitterService = twitterService
	_userService = userService
	_log = log
	_contentService = contentService
}

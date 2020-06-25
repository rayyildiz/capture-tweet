//go:generate go run github.com/99designs/gqlgen
package graph

import (
	"com.capturetweet/pkg/service"
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
func (r resolver) Subscription() SubscriptionResolver {
	return newSubscriptionResolver()
}

var (
	_twitterService service.TweetService = nil
	_userService    service.UserService  = nil
	_log            *zap.Logger          = nil
)

func InitService(log *zap.Logger, twitterService service.TweetService, userService service.UserService) {
	_twitterService = twitterService
	_userService = userService
	_log = log
}

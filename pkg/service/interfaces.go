//go:generate mockgen -package=service -self_package=com.capturetweet/pkg/service -destination=interfaces_mock.go . UserService,TweetService,SearchService
package service

import (
	"github.com/ChimeraCoder/anaconda"
)

type UserService interface {
	FindById(id string) (*UserModel, error)
	FindOrCreate(user *anaconda.User) (*UserModel, error)
}

type TweetService interface {
	FindById(id string) (*TweetModel, error)
	Store(url string) (string, error)
	Search(term string, size, start, page int) ([]TweetModel, error)
}

type SearchService interface {
	Search(term string, size int) ([]SearchModel, error)
	Put(tweetId, fullText, author string) error
}

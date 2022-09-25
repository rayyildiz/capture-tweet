//go:generate go run github.com/golang/mock/mockgen -package=infra -self_package=capturetweet.com/internal/infra -destination=anaconda_mock.go . TweetAPI
package infra

import (
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
)

func NewTwitterClient() TweetAPI {
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("TWITTER_ACCESS_SECRET")
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")

	return anaconda.NewTwitterApiWithCredentials(accessToken, accessTokenSecret, consumerKey, consumerSecret)
}

type TweetAPI interface {
	GetTweet(id int64, v url.Values) (tweet anaconda.Tweet, err error)
}

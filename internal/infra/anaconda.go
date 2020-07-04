//go:generate mockgen -package=infra -self_package=com.capturetweet/internal/infra -destination=anaconda_mock.go . TweetAPI
package infra

import (
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"os"
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

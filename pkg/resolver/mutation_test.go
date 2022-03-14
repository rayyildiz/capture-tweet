package resolver

import (
	"context"
	"github.com/matryer/is"
	"strings"
	"testing"

	"capturetweet.com/api"
	"capturetweet.com/internal/infra"
	"github.com/golang/mock/gomock"
)

func TestMutationResolver_Capture(t *testing.T) {
	is := is.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	infra.RegisterLogger()

	tweetService := api.NewMockTweetService(ctrl)
	tweetService.EXPECT().Store(gomock.Any(), "https://twitter.com/jack/status/20").Return("20", nil)
	tweetService.EXPECT().FindById(gomock.Any(), "20").Return(&api.TweetModel{
		ID:       "20",
		FullText: "full text",
		AuthorID: "jack",
	}, nil)

	_twitterService = tweetService
	resolver := newMutationResolver()

	tweet, err := resolver.Capture(context.Background(), "https://twitter.com/jack/status/20")
	is.NoErr(err)
	is.True(tweet != nil)

	is.True(tweet.ID == "20")
	is.True(tweet.AuthorID != nil)
	is.True(*tweet.AuthorID == "jack")
	is.True(tweet.FullText == "full text")

}

func TestMutationResolver_Contact(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	infra.RegisterLogger()

	contentService := api.NewMockContentService(ctrl)
	contentService.EXPECT().StoreContactRequest(gomock.Any(), "test@example.com", "Ramazan", "hello", "captcha").Return(nil)

	_contentService = contentService
	resolver := newMutationResolver()

	msg, err := resolver.Contact(context.Background(), ContactInput{
		FullName: "Ramazan",
		Email:    "test@example.com",
		Message:  "hello",
	}, nil, "captcha")

	is.NoErr(err)

	is.True(len(msg) > 0)
	is.True(strings.Contains(msg, "saved your message"))
}

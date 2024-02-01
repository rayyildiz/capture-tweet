package resolver

import (
	"context"
	"testing"

	"capturetweet.com/api"
	"capturetweet.com/internal/infra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestMutationResolver_Capture(t *testing.T) {
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
	require.NoError(t, err)
	if assert.NotNil(t, tweet) {
		assert.Equal(t, "20", tweet.ID)
		assert.Equal(t, "jack", *tweet.AuthorID)
		assert.Equal(t, "full text", tweet.FullText)
	}
}

func TestMutationResolver_Contact(t *testing.T) {
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

	require.NoError(t, err)
	if assert.True(t, len(msg) > 0) {
		assert.Regexp(t, "saved your message", msg)
	}

}

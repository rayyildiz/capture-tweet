package resolver

import (
	"com.capturetweet/api"
	"com.capturetweet/internal/infra"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMutationResolver_Capture(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	log := infra.NewLogger()
	require.NotNil(t, log)

	tweetService := api.NewMockTweetService(ctrl)
	tweetService.EXPECT().Store(ctx, "https://twitter.com/jack/status/20").Return("20", nil)
	tweetService.EXPECT().FindById(ctx, "20").Return(&api.TweetModel{
		ID:       "20",
		FullText: "full text",
		AuthorID: "jack",
	}, nil)

	_log = log
	_twitterService = tweetService
	resolver := newMutationResolver()

	tweet, err := resolver.Capture(ctx, "https://twitter.com/jack/status/20")
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
	ctx := context.Background()

	log := infra.NewLogger()
	require.NotNil(t, log)

	contentService := api.NewMockContentService(ctrl)
	contentService.EXPECT().StoreContactRequest(ctx, "test@example.com", "Ramazan", "hello", "captcha").Return(nil)

	_log = log
	_contentService = contentService
	resolver := newMutationResolver()

	msg, err := resolver.Contact(ctx, ContactInput{
		FullName: "Ramazan",
		Email:    "test@example.com",
		Message:  "hello",
	}, nil, "captcha")

	require.NoError(t, err)
	if assert.True(t, len(msg) > 0) {
		assert.Regexp(t, "saved your message", msg)
	}

}

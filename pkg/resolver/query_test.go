package resolver

import (
	"context"
	"testing"

	"capturetweet.com/api"
	"capturetweet.com/internal/infra"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueryResolver_Tweet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	infra.RegisterLogger()

	tweetService := api.NewMockTweetService(ctrl)
	tweetService.EXPECT().FindById(gomock.Any(), "1234").Return(&api.TweetModel{
		ID:       "1234",
		FullText: "test",
		PostedAt: nil,
	}, nil)

	_twitterService = tweetService
	resolver := newQueryResolver()

	tweet, err := resolver.Tweet(context.Background(), "1234")
	require.NoError(t, err)
	if assert.NotNil(t, tweet) {
		assert.Equal(t, "1234", tweet.ID)
		assert.Equal(t, "test", tweet.FullText)
	}
}

func TestQueryResolver_SearchByUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	infra.RegisterLogger()

	tweetService := api.NewMockTweetService(ctrl)
	tweetService.EXPECT().SearchByUser(gomock.Any(), "user1").Return([]api.TweetModel{
		{
			ID:            "1",
			FullText:      "test1",
			Lang:          "en",
			FavoriteCount: 5,
			RetweetCount:  10,
			AuthorID:      "user1",
		},
		{
			ID:            "2",
			FullText:      "test2",
			Lang:          "ar",
			FavoriteCount: 18,
			RetweetCount:  200,
			AuthorID:      "user1",
		},
	}, nil)

	_twitterService = tweetService
	resolver := newQueryResolver()

	tweets, err := resolver.SearchByUser(context.Background(), "user1")
	require.NoError(t, err)
	if assert.Equal(t, 2, len(tweets)) {
		assert.Equal(t, "1", tweets[0].ID)
		assert.Equal(t, "2", tweets[1].ID)

		if assert.NotNil(t, tweets[0].Lang) && assert.NotNil(t, tweets[1].Lang) {
			assert.Equal(t, "en", *tweets[0].Lang)
			assert.Equal(t, "ar", *tweets[1].Lang)
		}

		assert.Equal(t, "test1", tweets[0].FullText)
		assert.Equal(t, "test2", tweets[1].FullText)

		if assert.NotNil(t, tweets[0].AuthorID) && assert.NotNil(t, tweets[1].AuthorID) {
			assert.Equal(t, "user1", *tweets[0].AuthorID)
			assert.Equal(t, "user1", *tweets[1].AuthorID)
		}

		if assert.NotNil(t, tweets[0].FavoriteCount) && assert.NotNil(t, tweets[1].FavoriteCount) {
			assert.Equal(t, 5, *tweets[0].FavoriteCount)
			assert.Equal(t, 18, *tweets[1].FavoriteCount)
		}
		if assert.NotNil(t, tweets[0].RetweetCount) && assert.NotNil(t, tweets[1].RetweetCount) {
			assert.Equal(t, 10, *tweets[0].RetweetCount)
			assert.Equal(t, 200, *tweets[1].RetweetCount)
		}
	}
}

func TestQueryResolver_Search(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	infra.RegisterLogger()

	tweetService := api.NewMockTweetService(ctrl)
	tweetService.EXPECT().Search(gomock.Any(), "test", 10, 0, 0).Return([]api.TweetModel{
		{
			ID:            "1",
			FullText:      "test1",
			Lang:          "en",
			FavoriteCount: 5,
			RetweetCount:  10,
			AuthorID:      "user1",
		},
		{
			ID:            "2",
			FullText:      "test2",
			Lang:          "ar",
			FavoriteCount: 18,
			RetweetCount:  200,
			AuthorID:      "user2",
		},
		{
			ID:            "3",
			FullText:      "test3",
			Lang:          "tr",
			FavoriteCount: 0,
			RetweetCount:  0,
			AuthorID:      "user1",
		},
	}, nil)

	_twitterService = tweetService
	resolver := newQueryResolver()

	tweets, err := resolver.Search(context.Background(), SearchInput{Term: "test"}, 10, 0, 0)
	require.NoError(t, err)
	if assert.Equal(t, 3, len(tweets)) {
		assert.Equal(t, "1", tweets[0].ID)
		assert.Equal(t, "2", tweets[1].ID)
		assert.Equal(t, "3", tweets[2].ID)

		if assert.NotNil(t, tweets[0].Lang) && assert.NotNil(t, tweets[1].Lang) && assert.NotNil(t, tweets[2].Lang) {
			assert.Equal(t, "en", *tweets[0].Lang)
			assert.Equal(t, "ar", *tweets[1].Lang)
			assert.Equal(t, "tr", *tweets[2].Lang)
		}

		assert.Equal(t, "test1", tweets[0].FullText)
		assert.Equal(t, "test2", tweets[1].FullText)
		assert.Equal(t, "test3", tweets[2].FullText)

		if assert.NotNil(t, tweets[0].AuthorID) && assert.NotNil(t, tweets[1].AuthorID) && assert.NotNil(t, tweets[2].AuthorID) {
			assert.Equal(t, "user1", *tweets[0].AuthorID)
			assert.Equal(t, "user2", *tweets[1].AuthorID)
			assert.Equal(t, "user1", *tweets[2].AuthorID)
		}

		if assert.NotNil(t, tweets[0].FavoriteCount) && assert.NotNil(t, tweets[1].FavoriteCount) && assert.NotNil(t, tweets[2].FavoriteCount) {
			assert.Equal(t, 5, *tweets[0].FavoriteCount)
			assert.Equal(t, 18, *tweets[1].FavoriteCount)
			assert.Equal(t, 0, *tweets[2].FavoriteCount)
		}
		if assert.NotNil(t, tweets[0].RetweetCount) && assert.NotNil(t, tweets[1].RetweetCount) && assert.NotNil(t, tweets[2].RetweetCount) {

			assert.Equal(t, 10, *tweets[0].RetweetCount)
			assert.Equal(t, 200, *tweets[1].RetweetCount)
			assert.Equal(t, 0, *tweets[2].RetweetCount)
		}
	}
}

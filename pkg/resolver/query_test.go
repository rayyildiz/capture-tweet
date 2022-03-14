package resolver

import (
	"context"
	"github.com/matryer/is"
	"testing"

	"capturetweet.com/api"
	"capturetweet.com/internal/infra"
	"github.com/golang/mock/gomock"
)

func TestQueryResolver_Tweet(t *testing.T) {
	is := is.New(t)

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
	is.NoErr(err)
	is.NoErr(err)
	is.True(tweet != nil)

	is.Equal("1234", tweet.ID)
	is.Equal("test", tweet.FullText)
}

func TestQueryResolver_SearchByUser(t *testing.T) {
	is := is.New(t)

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
	is.NoErr(err)

	is.True(2 == len(tweets))

	is.Equal("1", tweets[0].ID)
	is.Equal("2", tweets[1].ID)

	is.True(tweets[0].Lang != nil)
	is.True(tweets[1].Lang != nil)
	is.Equal("en", *tweets[0].Lang)
	is.Equal("ar", *tweets[1].Lang)

	is.Equal("test1", tweets[0].FullText)
	is.Equal("test2", tweets[1].FullText)

	is.True(tweets[0].AuthorID != nil)
	is.True(tweets[1].AuthorID != nil)
	is.Equal("user1", *tweets[0].AuthorID)
	is.Equal("user1", *tweets[1].AuthorID)

	is.True(tweets[0].FavoriteCount != nil)
	is.True(tweets[1].FavoriteCount != nil)
	is.Equal(5, *tweets[0].FavoriteCount)
	is.Equal(18, *tweets[1].FavoriteCount)

	is.True(tweets[0].RetweetCount != nil)
	is.True(tweets[1].RetweetCount != nil)
	is.Equal(10, *tweets[0].RetweetCount)
	is.Equal(200, *tweets[1].RetweetCount)
}

func TestQueryResolver_Search(t *testing.T) {
	is := is.New(t)

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
	is.NoErr(err)

	is.Equal(3, len(tweets))
	is.Equal("1", tweets[0].ID)
	is.Equal("2", tweets[1].ID)
	is.Equal("3", tweets[2].ID)

	is.True(nil != tweets[0].Lang)
	is.True(nil != tweets[1].Lang)
	is.True(nil != tweets[2].Lang)
	is.Equal("en", *tweets[0].Lang)
	is.Equal("ar", *tweets[1].Lang)
	is.Equal("tr", *tweets[2].Lang)

	is.Equal("test1", tweets[0].FullText)
	is.Equal("test2", tweets[1].FullText)
	is.Equal("test3", tweets[2].FullText)

	is.True(nil != tweets[0].AuthorID)
	is.True(nil != tweets[1].AuthorID)
	is.True(nil != tweets[2].AuthorID)
	is.Equal("user1", *tweets[0].AuthorID)
	is.Equal("user2", *tweets[1].AuthorID)
	is.Equal("user1", *tweets[2].AuthorID)

	is.True(nil != tweets[0].FavoriteCount)
	is.True(nil != tweets[1].FavoriteCount)
	is.True(nil != tweets[2].FavoriteCount)
	is.Equal(5, *tweets[0].FavoriteCount)
	is.Equal(18, *tweets[1].FavoriteCount)
	is.Equal(0, *tweets[2].FavoriteCount)

	is.True(nil != tweets[0].RetweetCount)
	is.True(nil != tweets[1].RetweetCount)
	is.True(nil != tweets[2].RetweetCount)

	is.Equal(10, *tweets[0].RetweetCount)
	is.Equal(200, *tweets[1].RetweetCount)
	is.Equal(0, *tweets[2].RetweetCount)
}

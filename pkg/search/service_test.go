package search

import (
	"context"
	"github.com/matryer/is"
	"testing"

	"capturetweet.com/api"
	"capturetweet.com/internal/infra"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/golang/mock/gomock"
)

func TestService_Search(t *testing.T) {
	is := is.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	algolia := infra.NewMockIndexInterface(ctrl)
	algolia.EXPECT().Search("test", gomock.Any()).Return(search.QueryRes{
		Hits: []map[string]interface{}{
			{
				"objectID": "1",
				"fullText": "test",
				"author":   "rayyildiz",
			},
			{
				"objectID": "2",
				"fullText": "second tweet",
				"author":   "rayyildiz",
			},
		},
	}, nil)

	svc := NewService(algolia)

	searchModels, err := svc.Search(context.Background(), "test", 20)
	is.NoErr(err)
	is.True(searchModels != nil)

	is.True(len(searchModels) >= 2)

	is.Equal("1", searchModels[0].TweetID)
	is.Equal("test", searchModels[0].FullText)
	is.Equal("rayyildiz", searchModels[0].Author)

	is.Equal("2", searchModels[1].TweetID)
	is.Equal("second tweet", searchModels[1].FullText)
}

func TestService_Put(t *testing.T) {
	is := is.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	algolia := infra.NewMockIndexInterface(ctrl)
	algolia.EXPECT().SaveObject(api.SearchModel{
		TweetID:  "1",
		FullText: "test",
		Author:   "AYYILDIZ",
	}).Return(search.SaveObjectRes{}, nil)

	svc := NewService(algolia)

	err := svc.Put(context.Background(), "1", "test", "AYYILDIZ")
	is.NoErr(err)
}

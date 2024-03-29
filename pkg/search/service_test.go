package search

import (
	"context"
	"testing"

	"capturetweet.com/api"
	"capturetweet.com/internal/infra"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestService_Search(t *testing.T) {
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
	require.NoError(t, err)
	require.NotNil(t, searchModels)
	if assert.Equal(t, 2, len(searchModels)) {
		if assert.NotNil(t, searchModels[0]) {
			assert.Equal(t, "1", searchModels[0].TweetID)
			assert.Equal(t, "test", searchModels[0].FullText)
			assert.Equal(t, "rayyildiz", searchModels[0].Author)
		}

		if assert.NotNil(t, searchModels[1]) {
			assert.Equal(t, "2", searchModels[1].TweetID)
			assert.Equal(t, "second tweet", searchModels[1].FullText)
		}
	}
}

func TestService_Put(t *testing.T) {
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
	require.NoError(t, err)
}

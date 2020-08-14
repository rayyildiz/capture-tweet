package tweet

import (
	"com.capturetweet/api"
	"com.capturetweet/internal/infra"
	"context"
	"github.com/ChimeraCoder/anaconda"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestService_FindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	repo := NewMockRepository(ctrl)
	repo.EXPECT().FindById(ctx, "1").Return(&Tweet{
		ID:              "1",
		PostedAt:        time.Now(),
		FullText:        "text",
		CaptureURL:      nil,
		CaptureThumbURL: nil,
		AuthorID:        "user1",
		Resources: []Resource{
			{
				ID:        "img1",
				URL:       "https://path.com/example1.png",
				Width:     100,
				Height:    200,
				MediaType: "images/png",
			},
		},
	}, nil)
	log := infra.NewLogger()
	require.NotNil(t, log)

	svc := NewService(repo, nil, nil, nil, log, nil)
	require.NotNil(t, svc)

	model, err := svc.FindById(ctx, "1")
	require.NoError(t, err)
	if assert.NotNil(t, model) {
		assert.Equal(t, "1", model.ID)
		assert.Equal(t, "user1", model.AuthorID)
		assert.NotNil(t, model.PostedAt)
		assert.Nil(t, model.CaptureThumbURL)
		assert.Nil(t, model.CaptureURL)
		if assert.NotNil(t, model.Resources) {
			assert.Equal(t, 1, len(model.Resources))
			assert.Equal(t, "img1", model.Resources[0].ID)
			assert.Equal(t, "images/png", model.Resources[0].ResourceType)
			assert.Equal(t, 200, model.Resources[0].Height)
			assert.Equal(t, 100, model.Resources[0].Width)
		}
	}
}

func TestService_Store(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	repo := NewMockRepository(ctrl)
	repo.EXPECT().Exist(ctx, "20").Return(false)
	repo.EXPECT().Store(ctx, gomock.Any()).Return(nil)

	searchService := api.NewMockSearchService(ctrl)
	searchService.EXPECT().Put(ctx, "20", "test", "jack").Return(nil)

	userService := api.NewMockUserService(ctrl)
	userService.EXPECT().FindOrCreate(ctx, gomock.Any()).Return(nil, nil)

	twitterAPI := infra.NewMockTweetAPI(ctrl)
	twitterAPI.EXPECT().GetTweet(int64(20), gomock.Any()).Return(anaconda.Tweet{
		IdStr:    "20",
		FullText: "test",
		User: anaconda.User{
			ScreenName: "jack",
		},
	}, nil)

	topic, err := infra.NewTopic("mem://topicTest")
	require.NoError(t, err)

	log := infra.NewLogger()
	require.NotNil(t, log)

	svc := NewService(repo, searchService, userService, twitterAPI, log, topic)
	require.NotNil(t, svc)

	signal := make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		signal <- struct{}{}
	}()

	id, err := svc.Store(ctx, "https://twitter.com/jack/status/20")
	require.NoError(t, err)
	require.Equal(t, "20", id)
	<-signal
}

func TestService_UpdateLargeImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	repo := NewMockRepository(ctrl)
	repo.EXPECT().UpdateLargeImage(ctx, "1", "capture/large/1.png").Return(nil)
	log := infra.NewLogger()
	require.NotNil(t, log)

	svc := NewService(repo, nil, nil, nil, log, nil)
	require.NotNil(t, svc)

	err := svc.UpdateLargeImage(ctx, "1", "capture/large/1.png")
	require.NoError(t, err)
}

func TestService_UpdateThumbImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	repo := NewMockRepository(ctrl)
	repo.EXPECT().UpdateThumbImage(ctx, "2", "capture/thumb/2.png").Return(nil)
	log := infra.NewLogger()
	require.NotNil(t, log)

	svc := NewService(repo, nil, nil, nil, log, nil)
	require.NotNil(t, svc)

	err := svc.UpdateThumbImage(ctx, "2", "capture/thumb/2.png")
	require.NoError(t, err)
}

func TestService_SearchByUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	repo := NewMockRepository(ctrl)
	repo.EXPECT().FindByUser(ctx, "user1").Return([]Tweet{
		{
			ID:        "1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			PostedAt:  time.Now(),
			FullText:  "test1",
		},
		{
			ID:        "2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			PostedAt:  time.Now(),
			FullText:  "test2",
		},
	}, nil)

	log := infra.NewLogger()
	require.NotNil(t, log)

	svc := NewService(repo, nil, nil, nil, log, nil)
	require.NotNil(t, svc)

	tweets, err := svc.SearchByUser(ctx, "user1")
	require.NoError(t, err)
	if assert.Equal(t, 2, len(tweets)) {
		assert.Equal(t, "1", tweets[0].ID)
		assert.Equal(t, "2", tweets[1].ID)
		assert.Equal(t, "test1", tweets[0].FullText)
		assert.Equal(t, "test2", tweets[1].FullText)
	}
}

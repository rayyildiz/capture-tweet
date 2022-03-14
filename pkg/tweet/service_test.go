package tweet

import (
	"context"
	"github.com/matryer/is"
	"testing"
	"time"

	"capturetweet.com/api"
	"capturetweet.com/internal/infra"
	"github.com/ChimeraCoder/anaconda"
	"github.com/golang/mock/gomock"
)

func TestService_FindById(t *testing.T) {
	is := is.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockRepository(ctrl)
	repo.EXPECT().FindById(gomock.Any(), "1").Return(&Tweet{
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
	infra.RegisterLogger()

	svc := NewService(repo, nil, nil, nil, nil)
	is.True(svc != nil)

	model, err := svc.FindById(context.Background(), "1")
	is.NoErr(err)

	is.True(model != nil)

	is.Equal("1", model.ID)
	is.Equal("user1", model.AuthorID)

	is.True(nil != model.PostedAt)
	is.True(nil == model.CaptureThumbURL)
	is.True(nil == model.CaptureURL)
	is.True(nil != model.Resources)
	is.Equal(1, len(model.Resources))
	is.Equal("img1", model.Resources[0].ID)
	is.Equal("images/png", model.Resources[0].ResourceType)
	is.Equal(200, model.Resources[0].Height)
	is.Equal(100, model.Resources[0].Width)

}

func TestService_Store(t *testing.T) {
	is := is.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockRepository(ctrl)
	repo.EXPECT().Exist(gomock.Any(), "20").Return(false)
	repo.EXPECT().Store(gomock.Any(), gomock.Any()).Return(nil)

	searchService := api.NewMockSearchService(ctrl)
	searchService.EXPECT().Put(gomock.Any(), "20", "test", "jack").Return(nil)

	userService := api.NewMockUserService(ctrl)
	userService.EXPECT().FindOrCreate(gomock.Any(), gomock.Any()).Return(nil, nil)

	twitterAPI := infra.NewMockTweetAPI(ctrl)
	twitterAPI.EXPECT().GetTweet(int64(20), gomock.Any()).Return(anaconda.Tweet{
		IdStr:    "20",
		FullText: "test",
		User: anaconda.User{
			ScreenName: "jack",
		},
	}, nil)

	topic := infra.NewTopic("mem://topicTest")

	infra.RegisterLogger()

	svc := NewService(repo, searchService, userService, twitterAPI, topic)
	is.True(nil != svc)

	signal := make(chan struct{})
	go func() {
		time.Sleep(time.Second)
		signal <- struct{}{}
	}()

	id, err := svc.Store(context.Background(), "https://twitter.com/jack/status/20")
	is.NoErr(err)
	is.Equal("20", id)
	<-signal
}

func TestService_UpdateLargeImage(t *testing.T) {
	is := is.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	repo := NewMockRepository(ctrl)
	repo.EXPECT().UpdateLargeImage(gomock.Any(), "1", "capture/large/1.png").Return(nil)

	infra.RegisterLogger()

	svc := NewService(repo, nil, nil, nil, nil)
	is.True(nil != svc)

	err := svc.UpdateLargeImage(ctx, "1", "capture/large/1.png")
	is.NoErr(err)
}

func TestService_UpdateThumbImage(t *testing.T) {
	is := is.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockRepository(ctrl)
	repo.EXPECT().UpdateThumbImage(gomock.Any(), "2", "capture/thumb/2.png").Return(nil)

	infra.RegisterLogger()

	svc := NewService(repo, nil, nil, nil, nil)
	is.True(nil != svc)

	err := svc.UpdateThumbImage(context.Background(), "2", "capture/thumb/2.png")
	is.NoErr(err)
}

func TestService_SearchByUser(t *testing.T) {
	is := is.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockRepository(ctrl)
	repo.EXPECT().FindByUser(gomock.Any(), "user1").Return([]Tweet{
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

	infra.RegisterLogger()

	svc := NewService(repo, nil, nil, nil, nil)
	is.True(nil != svc)

	tweets, err := svc.SearchByUser(context.Background(), "user1")
	is.NoErr(err)
	is.Equal(2, len(tweets))

	is.Equal("1", tweets[0].ID)
	is.Equal("2", tweets[1].ID)
	is.Equal("test1", tweets[0].FullText)
	is.Equal("test2", tweets[1].FullText)
}

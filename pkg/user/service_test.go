package user

import (
	"context"
	"github.com/matryer/is"
	"testing"
	"time"

	"capturetweet.com/internal/infra"
	"github.com/ChimeraCoder/anaconda"
	"github.com/golang/mock/gomock"
)

func TestUserService_FindById(t *testing.T) {
	is := is.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	infra.RegisterLogger()

	ctx := context.Background()

	repo := NewMockRepository(ctrl)
	repo.EXPECT().FindById(gomock.Any(), "testUserIdStr").Return(&User{
		ID:         "testUserIdStr",
		CreatedAt:  time.Now(),
		RegisterAt: time.Now(),
		Username:   "rayyildiz",
		ScreenName: "Ramazan A.",
	}, nil)

	svc := NewService(repo)

	userModel, err := svc.FindById(ctx, "testUserIdStr")
	is.NoErr(err)
	is.True(nil != userModel)
	is.Equal("testUserIdStr", userModel.ID)
	is.Equal("rayyildiz", userModel.UserName)
	is.Equal("Ramazan A.", userModel.ScreenName)

}

func TestUserService_FindOrCreate_Exist(t *testing.T) {
	is := is.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	infra.RegisterLogger()

	ctx := context.Background()

	repo := NewMockRepository(ctrl)
	repo.EXPECT().FindById(gomock.Any(), "testId").Return(&User{
		ID:         "testId",
		CreatedAt:  time.Now(),
		RegisterAt: time.Now(),
		Username:   "rayyildiz",
		ScreenName: "Ramazan A.",
	}, nil)

	svc := NewService(repo)

	userModel, err := svc.FindOrCreate(ctx, &anaconda.User{
		IdStr:      "testId",
		ScreenName: "rayyildiz",
		Name:       "Ramazan A.",
	})
	is.NoErr(err)
	is.True(nil != userModel)
	is.Equal("testId", userModel.ID)
	is.Equal("rayyildiz", userModel.UserName)
	is.Equal("Ramazan A.", userModel.ScreenName)

}

func TestUserService_FindOrCreate_NotExist(t *testing.T) {
	is := is.New(t)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	infra.RegisterLogger()

	ctx := context.Background()

	repo := NewMockRepository(ctrl)

	strDate := "Mon Jan 02 15:04:05 -0700 2006"
	dt, err := time.Parse(time.RubyDate, strDate)
	is.NoErr(err)

	repo.EXPECT().FindById(gomock.Any(), "testId").Return(nil, nil)
	repo.EXPECT().Store(gomock.Any(), "testId", "rayyildiz", "Ramazan A.", "bio", "profile.png", dt).Return(nil)

	svc := NewService(repo)

	userModel, err := svc.FindOrCreate(ctx, &anaconda.User{
		IdStr:                "testId",
		ScreenName:           "rayyildiz",
		Name:                 "Ramazan A.",
		Description:          "bio",
		ProfileImageUrlHttps: "profile.png",
		CreatedAt:            strDate,
	})

	is.NoErr(err)
	is.True(nil != userModel)
	is.Equal("testId", userModel.ID)
	is.Equal("rayyildiz", userModel.UserName)
	is.Equal("Ramazan A.", userModel.ScreenName)

}

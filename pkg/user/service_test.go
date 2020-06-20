package user

import (
	"com.capturetweet/internal/infra"
	"github.com/ChimeraCoder/anaconda"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUserService_FindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	log := infra.NewLogger()
	require.NotNil(t, log)

	repo := NewMockRepository(ctrl)
	repo.EXPECT().FindById("testUserIdStr").Return(&User{
		ID:         "testUserIdStr",
		CreatedAt:  time.Now(),
		RegisterAt: time.Now(),
		Username:   "rayyildiz",
		ScreenName: "Ramazan A.",
	}, nil)

	svc := NewService(repo, log)

	userModel, err := svc.FindById("testUserIdStr")
	require.NoError(t, err)
	if assert.NotNil(t, userModel) {
		assert.Equal(t, "testUserIdStr", userModel.ID)
		assert.Equal(t, "rayyildiz", userModel.UserName)
		assert.Equal(t, "Ramazan A.", userModel.ScreenName)
	}
}

func TestUserService_FindOrCreate_Exist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	log := infra.NewLogger()
	require.NotNil(t, log)

	repo := NewMockRepository(ctrl)
	repo.EXPECT().FindById("testId").Return(&User{
		ID:         "testId",
		CreatedAt:  time.Now(),
		RegisterAt: time.Now(),
		Username:   "rayyildiz",
		ScreenName: "Ramazan A.",
	}, nil)

	svc := NewService(repo, log)

	userModel, err := svc.FindOrCreate(&anaconda.User{
		IdStr:      "testId",
		ScreenName: "rayyildiz",
		Name:       "Ramazan A.",
	})
	require.NoError(t, err)
	if assert.NotNil(t, userModel) {
		assert.Equal(t, "testId", userModel.ID)
		assert.Equal(t, "rayyildiz", userModel.UserName)
		assert.Equal(t, "Ramazan A.", userModel.ScreenName)
	}
}

func TestUserService_FindOrCreate_NotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	log := infra.NewLogger()
	require.NotNil(t, log)

	repo := NewMockRepository(ctrl)

	strDate := "Mon Jan 02 15:04:05 -0700 2006"
	dt, err := time.Parse(time.RubyDate, strDate)
	require.NoError(t, err)

	repo.EXPECT().FindById("testId").Return(nil, nil)
	repo.EXPECT().Store("testId", "rayyildiz", "Ramazan A.", "bio", "profile.png", dt).Return(nil)

	svc := NewService(repo, log)

	userModel, err := svc.FindOrCreate(&anaconda.User{
		IdStr:                "testId",
		ScreenName:           "rayyildiz",
		Name:                 "Ramazan A.",
		Description:          "bio",
		ProfileImageUrlHttps: "profile.png",
		CreatedAt:            strDate,
	})

	require.NoError(t, err)
	if assert.NotNil(t, userModel) {
		assert.Equal(t, "testId", userModel.ID)
		assert.Equal(t, "rayyildiz", userModel.UserName)
		assert.Equal(t, "Ramazan A.", userModel.ScreenName)
	}
}

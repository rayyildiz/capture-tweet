package user

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUserService_FindById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockRepository(ctrl)
	repo.EXPECT().FindById("testUserIdStr").Return(&User{
		ID:         "testUserIdStr",
		CreatedAt:  time.Now(),
		RegisterAt: time.Now(),
		Username:   "rayyildiz",
		ScreenName: "Ramazan A.",
	}, nil)

	svc := NewService(repo)

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

	repo := NewMockRepository(ctrl)
	repo.EXPECT().FindById("testId").Return(&User{
		ID:         "testId",
		CreatedAt:  time.Now(),
		RegisterAt: time.Now(),
		Username:   "rayyildiz",
		ScreenName: "Ramazan A.",
	}, nil)

	svc := NewService(repo)

	userModel, err := svc.FindOrCreate("testId", "rayyildiz", "Ramazan A.")
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

	repo := NewMockRepository(ctrl)

	repo.EXPECT().FindById("testId").Return(nil, nil)
	repo.EXPECT().Store("testId", "rayyildiz", "Ramazan A.").Return(nil)

	svc := NewService(repo)

	userModel, err := svc.FindOrCreate("testId", "rayyildiz", "Ramazan A.")
	require.NoError(t, err)
	if assert.NotNil(t, userModel) {
		assert.Equal(t, "testId", userModel.ID)
		assert.Equal(t, "rayyildiz", userModel.UserName)
		assert.Equal(t, "Ramazan A.", userModel.ScreenName)
	}
}

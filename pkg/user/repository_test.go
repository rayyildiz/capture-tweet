package user

import (
	"com.capturetweet/internal/infra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gocloud.dev/gcerrors"
	"testing"
	"time"
)

func TestRepository_Store(t *testing.T) {
	coll, err := infra.NewDocstore("mem://collection/id")
	require.NoError(t, err)
	defer coll.Close()

	repo := NewRepository(coll)

	err = repo.Store("testId", "username", "display name", "Bio", "profile.png", time.Now())
	require.NoError(t, err)
}

func TestRepository_FindById(t *testing.T) {
	coll, err := infra.NewDocstore("mem://collection2/id")
	require.NoError(t, err)
	defer coll.Close()

	repo := NewRepository(coll)

	id := "1270800178421706753"

	err = repo.Store(id, "test", "screenName", "Bio", "profile.png", time.Now())
	require.NoError(t, err)

	user, err := repo.FindById(id)
	if assert.NoError(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, "test", user.Username)
		assert.Equal(t, "screenName", user.ScreenName)
		assert.Equal(t, id, user.ID)
	}
}

func TestRepository_FindByUserName(t *testing.T) {
	coll, err := infra.NewDocstore("mem://collection3/id")
	require.NoError(t, err)
	defer coll.Close()

	repo := NewRepository(coll)

	require.NoError(t, repo.Store("1", "test1", "screenName 1", "Bio", "profile1.png", time.Now()))
	require.NoError(t, repo.Store("2", "test2", "screenName 2", "Bio", "profile2.png", time.Now()))
	require.NoError(t, repo.Store("3", "test3", "screenName 3", "Bio", "profile3.png", time.Now()))

	user, err := repo.FindByUserName("test3")
	if assert.NoError(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, "test3", user.Username)
		assert.Equal(t, "screenName 3", user.ScreenName)
		assert.Equal(t, "3", user.ID)
	}
}

func TestRepository_FindById_NotFound(t *testing.T) {
	coll, err := infra.NewDocstore("mem://collection4/id")
	require.NoError(t, err)
	defer coll.Close()

	repo := NewRepository(coll)

	user, err := repo.FindById("1")

	if assert.Nil(t, user) {
		assert.Error(t, err, "record not found")
		code := gcerrors.Code(err)
		assert.Equal(t, code, gcerrors.NotFound)
	}
}

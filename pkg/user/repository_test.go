package user

import (
	"com.capturetweet/pkg/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRepository_Store(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	require.NoError(t, err)
	defer db.Close()

	db.AutoMigrate(&model.User{})

	repo := NewRepository(db)

	err = repo.Store("testId", "username", "display name")
	require.NoError(t, err)
}

func TestRepository_FindById(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	require.NoError(t, err)
	defer db.Close()

	db.AutoMigrate(&model.User{})
	repo := NewRepository(db)

	id := "1270800178421706753"

	err = repo.Store(id, "test", "screenName")
	require.NoError(t, err)

	user, err := repo.FindById(id)
	if assert.NoError(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, "test", user.UserName)
		assert.Equal(t, "screenName", user.ScreenName)
		assert.Equal(t, id, user.ID)
	}
}

func TestMockRepository_FindByUserName(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	require.NoError(t, err)
	defer db.Close()

	db.AutoMigrate(&model.User{})
	repo := NewRepository(db)

	require.NoError(t, repo.Store("1", "test1", "screenName 1"))
	require.NoError(t, repo.Store("2", "test2", "screenName 2"))
	require.NoError(t, repo.Store("3", "test3", "screenName 3"))

	user, err := repo.FindByUserName("test3")
	if assert.NoError(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, "test3", user.UserName)
		assert.Equal(t, "screenName 3", user.ScreenName)
		assert.Equal(t, "3", user.ID)
	}
}

func TestRepository_FindById_NotFound(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	require.NoError(t, err)
	defer db.Close()

	db.AutoMigrate(&model.User{})
	repo := NewRepository(db)

	user, err := repo.FindById("1")

	if assert.Nil(t, user) {
		assert.Error(t, err, "record not found")
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	}
}

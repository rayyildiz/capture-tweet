package user

import (
	"capturetweet.com/internal/ent"
	"capturetweet.com/internal/ent/enttest"
	"context"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepository_Store(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	repo := NewRepository(client)

	err = repo.Store(context.Background(), "testId", "username", "display name", "Bio", "profile.png", time.Now())
	require.NoError(t, err)
}

func TestRepository_FindById(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	repo := NewRepository(client)

	id := "1270800178421706753"

	err = repo.Store(context.Background(), id, "test", "screenName", "Bio", "profile.png", time.Now())
	require.NoError(t, err)

	user, err := repo.FindById(context.Background(), id)
	if assert.NoError(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, "test", user.Username)
		assert.Equal(t, "screenName", user.ScreenName)
		assert.Equal(t, id, user.ID)
	}
}

func TestRepository_FindByUserName(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	repo := NewRepository(client)

	require.NoError(t, repo.Store(ctx, "1", "test1", "screenName 1", "Bio", "profile1.png", time.Now()))
	require.NoError(t, repo.Store(ctx, "2", "test2", "screenName 2", "Bio", "profile2.png", time.Now()))
	require.NoError(t, repo.Store(ctx, "3", "test3", "screenName 3", "Bio", "profile3.png", time.Now()))

	user, err := repo.FindByUserName(ctx, "test3")
	if assert.NoError(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, "test3", user.Username)
		assert.Equal(t, "screenName 3", user.ScreenName)
		assert.Equal(t, "3", user.ID)
	}
}

func TestRepository_FindById_NotFound(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	repo := NewRepository(client)

	user, err := repo.FindById(context.Background(), "1")

	if assert.Nil(t, user) {
		assert.Error(t, err, "record not found")
		notExist := ent.IsNotFound(err)
		assert.True(t, notExist)
	}
}

package user

import (
	"capturetweet.com/internal/ent"
	"capturetweet.com/internal/ent/enttest"
	"context"
	"github.com/matryer/is"
	"strings"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestRepository_Store(t *testing.T) {
	is := is.New(t)

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	is.NoErr(err)

	repo := NewRepository(client)

	err = repo.Store(context.Background(), "testId", "username", "display name", "Bio", "profile.png", time.Now())
	is.NoErr(err)
}

func TestRepository_FindById(t *testing.T) {
	is := is.New(t)

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	is.NoErr(err)

	repo := NewRepository(client)

	id := "1270800178421706753"

	err = repo.Store(context.Background(), id, "test", "screenName", "Bio", "profile.png", time.Now())
	is.NoErr(err)

	user, err := repo.FindById(context.Background(), id)

	is.NoErr(err)
	is.True(nil != user)
	is.Equal("test", user.Username)
	is.Equal("screenName", user.ScreenName)
	is.Equal(id, user.ID)

}

func TestRepository_FindByUserName(t *testing.T) {
	is := is.New(t)

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	is.NoErr(err)

	repo := NewRepository(client)

	is.NoErr(repo.Store(ctx, "1", "test1", "screenName 1", "Bio", "profile1.png", time.Now()))
	is.NoErr(repo.Store(ctx, "2", "test2", "screenName 2", "Bio", "profile2.png", time.Now()))
	is.NoErr(repo.Store(ctx, "3", "test3", "screenName 3", "Bio", "profile3.png", time.Now()))

	user, err := repo.FindByUserName(ctx, "test3")

	is.NoErr(err)
	is.True(nil != user)
	is.Equal("test3", user.Username)
	is.Equal("screenName 3", user.ScreenName)
	is.Equal("3", user.ID)

}

func TestRepository_FindById_NotFound(t *testing.T) {
	is := is.New(t)

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	is.NoErr(err)

	repo := NewRepository(client)

	user, err := repo.FindById(context.Background(), "1")

	is.True(nil == user)
	a := err.Error()
	is.True(strings.ContainsAny(a, "record not found"))
	notExist := ent.IsNotFound(err)
	is.True(notExist)

}

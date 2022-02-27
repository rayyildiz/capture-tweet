package tweet

import (
	"capturetweet.com/internal/ent/enttest"
	"context"
	"testing"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

func TestRepository_Store(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	repo := NewRepository(client)
	_, err = client.User.Create().SetID("1").SetUsername("rayyildiz").SetScreenName("Ramazan").SetRegisteredAt(time.Now()).Save(ctx)
	require.NoError(t, err)
	err = repo.Store(ctx, &anaconda.Tweet{
		IdStr:         "1",
		RetweetCount:  4,
		FavoriteCount: 9,
		Lang:          "en",
		FullText:      "hello",
		User: anaconda.User{
			IdStr: "1",
		},
		CreatedAt: time.Now().Format(time.RubyDate),
	})
	require.NoError(t, err)
}

func TestRepository_Exist(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	repo := NewRepository(client)
	_, err = client.User.Create().SetID("1").SetUsername("rayyildiz").SetScreenName("Ramazan").SetRegisteredAt(time.Now()).Save(ctx)
	require.NoError(t, err)
	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "testId1", FullText: "hello", Lang: "en", CreatedAt: time.Now().Format(time.RubyDate), User: anaconda.User{IdStr: "1"}})
	require.NoError(t, err)

	b := repo.Exist(ctx, "testId1")
	require.True(t, b)
}

func TestRepository_FindByIds(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	repo := NewRepository(client)

	_, err = client.User.Create().SetID("1").SetUsername("rayyildiz").SetScreenName("Ramazan").SetRegisteredAt(time.Now()).Save(ctx)
	require.NoError(t, err)

	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "1", FullText: "1", Lang: "en", CreatedAt: time.Now().Format(time.RubyDate), User: anaconda.User{IdStr: "1"}})
	require.NoError(t, err)

	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "2", FullText: "2", Lang: "en", CreatedAt: time.Now().Format(time.RubyDate), User: anaconda.User{IdStr: "1"}})
	require.NoError(t, err)

	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "3", FullText: "3", Lang: "en", CreatedAt: time.Now().Format(time.RubyDate), User: anaconda.User{IdStr: "1"}})
	require.NoError(t, err)

	tweets, err := repo.FindByIds(ctx, []string{"1", "3", "4"})
	require.NoError(t, err)
	if assert.Equal(t, 2, len(tweets)) {
		assert.Equal(t, "1", tweets[0].ID)
		assert.Equal(t, "3", tweets[1].ID)
	}
}

func TestRepository_UpdateThumbImage(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	repo := NewRepository(client)

	_, err = client.User.Create().SetID("1").SetUsername("rayyildiz").SetScreenName("Ramazan").SetRegisteredAt(time.Now()).Save(ctx)
	require.NoError(t, err)

	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "1", Lang: "en", FullText: "1", CreatedAt: time.Now().Format(time.RubyDate), User: anaconda.User{IdStr: "1"}})
	require.NoError(t, err)

	err = repo.UpdateThumbImage(ctx, "1", "image1.png")
	require.NoError(t, err)

	tweet, err := repo.FindById(ctx, "1")
	require.NoError(t, err)
	if assert.NotNil(t, tweet) {
		assert.Equal(t, "1", tweet.ID)
		if assert.NotNil(t, tweet.CaptureThumbURL) {
			assert.Equal(t, "image1.png", *tweet.CaptureThumbURL)
		}
	}
}

func TestRepository_FindByUser(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	repo := NewRepository(client)
	_, err = client.User.Create().SetID("1").SetUsername("rayyildiz").SetScreenName("Ramazan").SetRegisteredAt(time.Now()).Save(ctx)
	require.NoError(t, err)

	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "x1", Lang: "en", FullText: "1", CreatedAt: time.Now().Format(time.RubyDate), User: anaconda.User{IdStr: "1"}})
	require.NoError(t, err)

	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "ay2", Lang: "en", FullText: "2", CreatedAt: time.Now().Format(time.RubyDate), User: anaconda.User{IdStr: "1"}})
	require.NoError(t, err)

	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "z3", Lang: "en", FullText: "3", CreatedAt: time.Now().Format(time.RubyDate), User: anaconda.User{IdStr: "1"}})
	require.NoError(t, err)

	tweets, err := repo.FindByUser(ctx, "1")
	require.NoError(t, err)
	if assert.Equal(t, 3, len(tweets)) {
		ids := []string{tweets[0].ID, tweets[1].ID, tweets[2].ID}
		assert.Contains(t, ids, "x1")
		assert.Contains(t, ids, "ay2")
		assert.Contains(t, ids, "z3")
	}
}

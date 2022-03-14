package tweet

import (
	"capturetweet.com/internal/ent/enttest"
	"context"
	"github.com/matryer/is"
	"testing"
	"time"

	"github.com/ChimeraCoder/anaconda"
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
	_, err = client.User.Create().SetID("1").SetUsername("rayyildiz").SetScreenName("Ramazan").SetRegisteredAt(time.Now()).Save(ctx)
	is.NoErr(err)
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
	is.NoErr(err)
}

func TestRepository_Exist(t *testing.T) {
	is := is.New(t)

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	is.NoErr(err)

	repo := NewRepository(client)
	_, err = client.User.Create().SetID("1").SetUsername("rayyildiz").SetScreenName("Ramazan").SetRegisteredAt(time.Now()).Save(ctx)
	is.NoErr(err)
	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "testId1", FullText: "hello", Lang: "en", CreatedAt: time.Now().Format(time.RubyDate), User: anaconda.User{IdStr: "1"}})
	is.NoErr(err)

	b := repo.Exist(ctx, "testId1")
	is.True(b)
}

func TestRepository_FindByIds(t *testing.T) {
	is := is.New(t)

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	is.NoErr(err)

	repo := NewRepository(client)

	_, err = client.User.Create().SetID("1").SetUsername("rayyildiz").SetScreenName("Ramazan").SetRegisteredAt(time.Now()).Save(ctx)
	is.NoErr(err)

	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "1", FullText: "1", Lang: "en", CreatedAt: time.Now().Format(time.RubyDate), User: anaconda.User{IdStr: "1"}})
	is.NoErr(err)

	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "2", FullText: "2", Lang: "en", CreatedAt: time.Now().Format(time.RubyDate), User: anaconda.User{IdStr: "1"}})
	is.NoErr(err)

	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "3", FullText: "3", Lang: "en", CreatedAt: time.Now().Format(time.RubyDate), User: anaconda.User{IdStr: "1"}})
	is.NoErr(err)

	tweets, err := repo.FindByIds(ctx, []string{"1", "3", "4"})
	is.NoErr(err)

	is.Equal(2, len(tweets))
	is.Equal("1", tweets[0].ID)
	is.Equal("3", tweets[1].ID)
}

func TestRepository_UpdateThumbImage(t *testing.T) {
	is := is.New(t)

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	is.NoErr(err)

	repo := NewRepository(client)

	_, err = client.User.Create().SetID("1").SetUsername("rayyildiz").SetScreenName("Ramazan").SetRegisteredAt(time.Now()).Save(ctx)
	is.NoErr(err)

	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "1", Lang: "en", FullText: "1", CreatedAt: time.Now().Format(time.RubyDate), User: anaconda.User{IdStr: "1"}})
	is.NoErr(err)

	err = repo.UpdateThumbImage(ctx, "1", "image1.png")
	is.NoErr(err)

	tweet, err := repo.FindById(ctx, "1")
	is.NoErr(err)

	is.True(tweet != nil)

	is.Equal("1", tweet.ID)
	is.True(tweet.CaptureThumbURL != nil)
	is.Equal("image1.png", *tweet.CaptureThumbURL)
}

func TestRepository_FindByUser(t *testing.T) {
	is := is.New(t)

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	is.NoErr(err)

	repo := NewRepository(client)
	_, err = client.User.Create().SetID("1").SetUsername("rayyildiz").SetScreenName("Ramazan").SetRegisteredAt(time.Now()).Save(ctx)
	is.NoErr(err)

	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "x1", Lang: "en", FullText: "1", CreatedAt: time.Now().Format(time.RubyDate), User: anaconda.User{IdStr: "1"}})
	is.NoErr(err)

	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "ay2", Lang: "en", FullText: "2", CreatedAt: time.Now().Format(time.RubyDate), User: anaconda.User{IdStr: "1"}})
	is.NoErr(err)

	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "z3", Lang: "en", FullText: "3", CreatedAt: time.Now().Format(time.RubyDate), User: anaconda.User{IdStr: "1"}})
	is.NoErr(err)

	tweets, err := repo.FindByUser(ctx, "1")
	is.NoErr(err)

	is.Equal(3, len(tweets))

	ids := []string{tweets[0].ID, tweets[1].ID, tweets[2].ID}

	is.Equal(ids[0], "x1")
	is.Equal(ids[1], "ay2")
	is.Equal(ids[2], "z3")
}

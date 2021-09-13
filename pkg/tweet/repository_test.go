package tweet

import (
	"com.capturetweet/internal/infra"
	"context"
	"github.com/ChimeraCoder/anaconda"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRepository_Store(t *testing.T) {
	coll := infra.NewDocstore("mem://collection/id")
	defer coll.Close()

	ctx := context.Background()

	repo := NewRepository(coll)
	err := repo.Store(ctx, &anaconda.Tweet{
		IdStr:         "1",
		RetweetCount:  4,
		FavoriteCount: 9,
		Lang:          "en",
		CreatedAt:     time.Now().Format(time.RubyDate),
	})
	require.NoError(t, err)
}

func TestRepository_Exist(t *testing.T) {
	coll := infra.NewDocstore("mem://collection/id")
	defer coll.Close()
	ctx := context.Background()

	repo := NewRepository(coll)
	err := repo.Store(ctx, &anaconda.Tweet{IdStr: "testId1", Lang: "en", CreatedAt: time.Now().Format(time.RubyDate)})
	require.NoError(t, err)

	b := repo.Exist(ctx, "testId1")
	require.True(t, b)
}

func TestRepository_FindByIds(t *testing.T) {
	coll := infra.NewDocstore("mem://collection/id")
	defer coll.Close()

	ctx := context.Background()

	repo := NewRepository(coll)
	err := repo.Store(ctx, &anaconda.Tweet{IdStr: "1", Lang: "en", CreatedAt: time.Now().Format(time.RubyDate)})
	require.NoError(t, err)

	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "2", Lang: "en", CreatedAt: time.Now().Format(time.RubyDate)})
	require.NoError(t, err)

	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "3", Lang: "en", CreatedAt: time.Now().Format(time.RubyDate)})
	require.NoError(t, err)

	tweets, err := repo.FindByIds(ctx, []string{"1", "3", "4"})
	require.NoError(t, err)
	if assert.Equal(t, 2, len(tweets)) {
		assert.Equal(t, "1", tweets[0].ID)
		assert.Equal(t, "3", tweets[1].ID)
	}
}

func TestRepository_UpdateThumbImage(t *testing.T) {
	coll := infra.NewDocstore("mem://collection/id")
	defer coll.Close()

	ctx := context.Background()

	repo := NewRepository(coll)
	err := repo.Store(ctx, &anaconda.Tweet{IdStr: "1", Lang: "en", CreatedAt: time.Now().Format(time.RubyDate)})
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
	coll := infra.NewDocstore("mem://collection/id")
	defer coll.Close()

	ctx := context.Background()

	repo := NewRepository(coll)

	err := repo.Store(ctx, &anaconda.Tweet{IdStr: "x1", Lang: "en", CreatedAt: time.Now().Format(time.RubyDate), User: anaconda.User{IdStr: "123"}})
	require.NoError(t, err)

	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "ay2", Lang: "en", CreatedAt: time.Now().Format(time.RubyDate), User: anaconda.User{IdStr: "123"}})
	require.NoError(t, err)

	err = repo.Store(ctx, &anaconda.Tweet{IdStr: "z3", Lang: "en", CreatedAt: time.Now().Format(time.RubyDate), User: anaconda.User{IdStr: "123"}})
	require.NoError(t, err)

	tweets, err := repo.FindByUser(ctx, "123")
	require.NoError(t, err)
	if assert.Equal(t, 3, len(tweets)) {
		ids := []string{tweets[0].ID, tweets[1].ID, tweets[2].ID}
		assert.Contains(t, ids, "x1")
		assert.Contains(t, ids, "ay2")
		assert.Contains(t, ids, "z3")
	}
}

package tweet

import (
	"com.capturetweet/internal/infra"
	"github.com/ChimeraCoder/anaconda"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRepository_Store(t *testing.T) {
	coll, err := infra.NewDocstore("mem://collection/id")
	require.NoError(t, err)
	defer coll.Close()

	repo := NewRepository(coll)
	err = repo.Store(&anaconda.Tweet{
		IdStr: "1",
		User: anaconda.User{
			IdStr:      "1",
			ScreenName: "rayyildiz",
		},
		Lang:      "en",
		CreatedAt: time.Now().String(),
	})
	require.NoError(t, err)
}

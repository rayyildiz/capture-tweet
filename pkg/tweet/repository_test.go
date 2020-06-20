package tweet

import (
	"com.capturetweet/internal/infra"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRepository_Store(t *testing.T) {
	coll, err := infra.NewDocstore("mem://collection/id")
	require.NoError(t, err)
	defer coll.Close()

	repo := NewRepository(coll)
	err = repo.Store("id1", "full text", "en", "userId1", 1, 1, nil)
	require.NoError(t, err)
}

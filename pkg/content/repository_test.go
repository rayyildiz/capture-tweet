package content

import (
	"com.capturetweet/internal/infra"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRepository_ContactUS(t *testing.T) {
	coll, err := infra.NewDocstore("mem://collection/id")
	require.NoError(t, err)
	defer coll.Close()

	repo := NewRepository(coll)

	err = repo.ContactUs(context.Background(), "test", "ramazan", "hello")
	require.NoError(t, err)
}

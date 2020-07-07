package content

import (
	"com.capturetweet/internal/infra"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRepository_ContactUS(t *testing.T) {
	coll, err := infra.NewDocstore("mem://collection/id")
	require.NoError(t, err)
	defer coll.Close()

	repo := NewRepository(coll)

	err = repo.ContactUs("test", "ramazan", "hello")
	require.NoError(t, err)
}
package content

import (
	"context"
	"testing"

	"capturetweet.com/internal/infra"
	"github.com/stretchr/testify/require"
)

func TestRepository_ContactUS(t *testing.T) {
	coll := infra.NewDocstore("mem://collection/id")
	defer coll.Close()

	repo := NewRepository(coll)

	err := repo.ContactUs(context.Background(), "test", "ramazan", "hello")
	require.NoError(t, err)
}

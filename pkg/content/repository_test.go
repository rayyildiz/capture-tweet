package content

import (
	"capturetweet.com/internal/ent/enttest"
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

func TestRepository_ContactUS(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	repo := NewRepository(client)

	err = repo.ContactUs(context.Background(), "test", "ramazan", "hello")
	require.NoError(t, err)
}

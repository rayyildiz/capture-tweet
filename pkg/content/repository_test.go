package content

import (
	"capturetweet.com/internal/infra/database"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepository_ContactUS(t *testing.T) {
	db := database.NewInmemoryDb()
	defer db.Close()

	_, err := db.Exec(`CREATE TABLE contact_us
(
    id         text primary key,
    email      text,
    full_name  text,
    message    text,
    created_at timestamp
)`)
	require.NoError(t, err)

	repo := NewRepository(db)

	err = repo.ContactUs(context.Background(), "test", "ramazan", "hello")
	require.NoError(t, err)
}

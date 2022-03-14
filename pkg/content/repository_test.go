package content

import (
	"capturetweet.com/internal/ent/enttest"
	"context"
	"github.com/matryer/is"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestRepository_ContactUS(t *testing.T) {
	is := is.New(t)

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	ctx := context.Background()
	err := client.Schema.Create(ctx)
	is.NoErr(err)

	repo := NewRepository(client)

	err = repo.ContactUs(context.Background(), "test", "ramazan", "hello")
	is.NoErr(err)
}

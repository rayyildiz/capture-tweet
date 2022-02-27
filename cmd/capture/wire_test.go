package main

import (
	"os"
	"testing"
)

func TestInitializeBrowserService(t *testing.T) {
	os.Setenv("POSTGRES_CONNECTION", "host=localhost port=5432 user=postgres password=postgres")
	os.Setenv("BLOB_BUCKET", "mem://file")

	s := initializeBrowserService()
	if s == nil {
		t.Fatal("could not initialize wire service")
	}
}

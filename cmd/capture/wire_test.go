package main

import (
	"os"
	"testing"
)

func TestInitializeBrowserService(t *testing.T) {
	os.Setenv("DOCSTORE_TWEETS", "mem://tweets/id")
	os.Setenv("BLOB_BUCKET", "mem://file")

	s := initializeBrowserService()
	if s == nil {
		t.Fatal("could not initialize wire service")
	}
}

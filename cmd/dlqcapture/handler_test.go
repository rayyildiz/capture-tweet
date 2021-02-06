package main

import (
	"os"
	"testing"
)

func TestHandleDlq(t *testing.T) {
	os.Setenv("BLOB_BUCKET", "mem://bucket_dql2")

}

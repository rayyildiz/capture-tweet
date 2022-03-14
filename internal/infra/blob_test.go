package infra

import (
	"context"
	"github.com/matryer/is"
	"os"
	"testing"
)

func TestNewBucketFromEnvironment(t *testing.T) {
	is := is.New(t)

	os.Setenv("BLOB_BUCKET", "mem://bucket/to/memory")
	bucket := NewBucketFromEnvironment()
	defer bucket.Close()
	is.True(nil != bucket)

	err := bucket.WriteAll(context.Background(), "test", []byte("hello world"), nil)
	is.NoErr(err)
}

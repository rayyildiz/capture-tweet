package infra

import (
	"context"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestNewBucketFromEnvironment(t *testing.T) {
	os.Setenv("BLOB_BUCKET", "mem://bucket/to/memory")
	bucket := NewBucketFromEnvironment()
	defer bucket.Close()
	require.NotNil(t, bucket)

	err := bucket.WriteAll(context.Background(), "test", []byte("hello world"), nil)
	require.NoError(t, err)
}

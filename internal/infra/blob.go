package infra

import (
	"context"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/memblob"
	"os"
)

func NewBucket(bucketName string) (*blob.Bucket, error) {
	return blob.OpenBucket(context.Background(), bucketName)
}

func NewBucketFromEnvironment() (*blob.Bucket, error) {
	return NewBucket(os.Getenv("BLOB_BUCKET"))
}

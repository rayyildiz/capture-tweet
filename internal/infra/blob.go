package infra

import (
	"context"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/memblob"
	"log"
	"os"
)

func NewBucket(bucketName string) *blob.Bucket {
	bucket, err := blob.OpenBucket(context.Background(), bucketName)
	if err != nil {
		log.Fatalf("could not connect to blob storage, %v", err)
	}
	return bucket
}

func NewBucketFromEnvironment() *blob.Bucket {
	return NewBucket(os.Getenv("BLOB_BUCKET"))
}

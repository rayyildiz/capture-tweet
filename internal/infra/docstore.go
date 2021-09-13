package infra

import (
	"context"
	"log"
	"os"

	"gocloud.dev/docstore"
	_ "gocloud.dev/docstore/gcpfirestore"
	_ "gocloud.dev/docstore/memdocstore"
)

func NewDocstore(collectionName string) *docstore.Collection {
	if len(collectionName) < 1 {
		log.Fatalf("collection name is empty")
	}

	coll, err := docstore.OpenCollection(context.Background(), collectionName)
	if err != nil {
		log.Fatalf("error while opening docstore, %v", err)
	}
	return coll
}

func NewTweetCollection() *docstore.Collection {
	return NewDocstore(os.Getenv("DOCSTORE_TWEETS"))
}

func NewUserCollection() *docstore.Collection {
	return NewDocstore(os.Getenv("DOCSTORE_USERS"))
}

func NewContactUsCollection() *docstore.Collection {
	return NewDocstore(os.Getenv("DOCSTORE_CONTACT_US"))
}

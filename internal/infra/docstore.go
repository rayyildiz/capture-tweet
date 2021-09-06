package infra

import (
	"context"
	"errors"
	"os"

	"gocloud.dev/docstore"
	_ "gocloud.dev/docstore/gcpfirestore"
	_ "gocloud.dev/docstore/memdocstore"
)

var (
	ErrInvalidDocstoreEnv = errors.New("invalid collection name")
)

func NewDocstore(collectionName string) (*docstore.Collection, error) {
	if len(collectionName) < 1 {
		return nil, ErrInvalidDocstoreEnv
	}

	return docstore.OpenCollection(context.Background(), collectionName)
}

func NewTweetCollection() (*docstore.Collection, error) {
	return NewDocstore(os.Getenv("DOCSTORE_TWEETS"))
}

func NewUserCollection() (*docstore.Collection, error) {
	return NewDocstore(os.Getenv("DOCSTORE_USERS"))
}

func NewContactUsCollection() (*docstore.Collection, error) {
	return NewDocstore(os.Getenv("DOCSTORE_CONTACT_US"))
}

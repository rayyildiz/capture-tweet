package infra

import (
	"context"
	"github.com/matryer/is"
	"os"
	"testing"
	"time"
)

func TestNewContactUsCollection(t *testing.T) {
	is := is.New(t)

	os.Setenv("DOCSTORE_CONTACT_US", "mem://contact/ID")

	coll := NewContactUsCollection()
	is.True(nil != coll)

	err := coll.Put(context.Background(), map[string]interface{}{
		"ID":      1,
		"Email":   "email",
		"Subject": "subject",
		"Date":    time.Now(),
	})
	is.NoErr(err)
}

func TestNewTweetCollection(t *testing.T) {
	is := is.New(t)

	os.Setenv("DOCSTORE_TWEETS", "mem://tweet/id")

	coll := NewTweetCollection()
	is.True(nil != coll)

	err := coll.Put(context.Background(), map[string]interface{}{
		"id":         1,
		"user":       "user",
		"text":       "test",
		"created_at": time.Now(),
	})
	is.NoErr(err)
}

func TestNewUserCollection(t *testing.T) {
	is := is.New(t)

	os.Setenv("DOCSTORE_USERS", "mem://users/Username")

	coll := NewUserCollection()
	is.True(nil != coll)

	err := coll.Put(context.Background(), map[string]interface{}{
		"Username":  "@rayyildiz",
		"FullName":  "Ramazan",
		"CreatedAt": time.Now(),
	})
	is.NoErr(err)
}

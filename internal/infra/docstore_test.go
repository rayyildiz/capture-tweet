package infra

import (
	"context"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestNewContactUsCollection(t *testing.T) {
	os.Setenv("DOCSTORE_CONTACT_US", "mem://contact/ID")

	coll, err := NewContactUsCollection()
	require.NoError(t, err)
	require.NotNil(t, coll)

	err = coll.Put(context.Background(), map[string]interface{}{
		"ID":      1,
		"Email":   "email",
		"Subject": "subject",
		"Date":    time.Now(),
	})
	require.NoError(t, err)
}

func TestNewTweetCollection(t *testing.T) {
	os.Setenv("DOCSTORE_TWEETS", "mem://tweet/id")

	coll, err := NewTweetCollection()
	require.NoError(t, err)
	require.NotNil(t, coll)

	err = coll.Put(context.Background(), map[string]interface{}{
		"id":         1,
		"user":       "user",
		"text":       "test",
		"created_at": time.Now(),
	})
	require.NoError(t, err)
}

func TestNewUserCollection(t *testing.T) {
	os.Setenv("DOCSTORE_USERS", "mem://users/Username")

	coll, err := NewUserCollection()
	require.NoError(t, err)
	require.NotNil(t, coll)

	err = coll.Put(context.Background(), map[string]interface{}{
		"Username":  "@rayyildiz",
		"FullName":  "Ramazan",
		"CreatedAt": time.Now(),
	})
	require.NoError(t, err)
}

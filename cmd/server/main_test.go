package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	// defer goleak.VerifyNone(t)

	rand.Seed(time.Now().Unix())
	port := rand.Intn(2000) + 30010
	log.Printf("port is %d", port)

	os.Setenv("DOCSTORE_TWEETS", "mem://tweet/ID")
	os.Setenv("BLOB_BUCKET", "mem://bucket")
	os.Setenv("DOCSTORE_USERS", "mem://users/ID")
	os.Setenv("DOCSTORE_CONTACT_US", "mem://contact/ID")
	os.Setenv("TOPIC_CAPTURE", "mem://topic1")
	os.Setenv("ALGOLIA_SECRET", "secret")
	os.Setenv("ALGOLIA_CLIENT_ID", "client")
	os.Setenv("ALGOLIA_INDEX", "test")
	os.Setenv("PORT", strconv.Itoa(port))

	signal := make(chan struct{})

	go func() {
		log.Printf("sleep 2 seconds")
		time.Sleep(2 * time.Second)
		log.Printf("close the app")
		signal <- struct{}{}
	}()

	go func() {
		err := Run()
		require.NoError(t, err)
	}()

	<-signal
}

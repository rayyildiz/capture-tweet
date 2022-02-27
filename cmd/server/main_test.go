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
	//defer goleak.VerifyNone(t)

	rand.Seed(time.Now().Unix())
	port := rand.Intn(2000) + 30010
	log.Printf("port is %d", port)

	os.Setenv("POSTGRES_CONNECTION", "host=localhost port=5432 user=postgres password=postgres")
	os.Setenv("BLOB_BUCKET", "mem://bucket")
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

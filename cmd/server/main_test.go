package main

import (
	"github.com/matryer/is"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	is := is.New(t)

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
		is.NoErr(err)
	}()

	<-signal
}

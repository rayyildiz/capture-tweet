package main

import (
	"github.com/stretchr/testify/require"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	rand.Seed(time.Now().Unix())
	port := rand.Intn(2000) + 20700
	log.Printf("port is %d", port)

	os.Setenv("BLOB_BUCKET", "mem://bucket_dlq")
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

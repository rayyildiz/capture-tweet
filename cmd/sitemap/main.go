package main

import (
	"com.capturetweet/internal/infra"
	"com.capturetweet/pkg/tweet"
	"github.com/joho/godotenv"
	. "go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	godotenv.Load()
}

func main() {
	start := time.Now()

	port := os.Getenv("PORT")
	if port == "" {
		port = "4400"
	}

	logger := infra.NewLogger()
	ensureNotNil(logger, "zap:logger")

	tweetColl, err := infra.NewTweetCollection()
	ensureNoError(err, "twitter:docstore collection")
	defer tweetColl.Close()

	h := handlerImpl{
		log:  logger,
		repo: tweet.NewRepository(tweetColl),
	}

	http.HandleFunc("/sitemap", h.handleRequest)

	diff := time.Now().Sub(start)
	logger.Info("initialized objects", Duration("elapsed", diff))

	err = http.ListenAndServe(":"+port, nil)
	ensureNoError(err, "http:ListenAndServe, port :"+port)
}

func ensureNoError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s, %v", msg, err)
	}
}

func ensureNotNil(obj interface{}, msg string) {
	if obj == nil {
		log.Fatalf("object is nil, %s", msg)
	}
}

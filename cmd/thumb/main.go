package main

import (
	"com.capturetweet/internal/infra"
	"com.capturetweet/pkg/tweet"
	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	godotenv.Load()
}

func main() {
	err := infra.InitSentry()
	ensureNoError(err, "sentry init")
	defer sentry.Flush(time.Second * 2)

	start := time.Now()
	logger := infra.NewLogger()
	ensureNotNil(logger, "zap:logger")

	tweetColl, err := infra.NewTweetCollection()
	ensureNoError(err, "twitter:docstore collection")
	defer tweetColl.Close()

	bucket, err := infra.NewBucketFromEnvironment()
	ensureNoError(err, "blob bucket")
	defer bucket.Close()

	tweetService := tweet.NewService(tweet.NewRepository(tweetColl), nil, nil, nil, logger, nil)
	ensureNotNil(tweetService, "tweet service initialize")

	port := os.Getenv("PORT")
	if port == "" {
		port = "4300"
	}

	h := handlerImpl{
		log:     logger,
		service: tweetService,
		bucket:  bucket,
	}
	http.HandleFunc("/resize", h.handleResize)

	diff := time.Now().Sub(start)
	logger.Info("initialized objects", zap.Duration("elapsed", diff))

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

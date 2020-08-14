package main

import (
	"com.capturetweet/internal/infra"
	"com.capturetweet/pkg/tweet"
	"fmt"
	"github.com/getsentry/sentry-go"
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
	if err := Run(); err != nil {
		log.Printf("%v", err)
		os.Exit(1)
	}
}

func Run() error {
	err := infra.InitSentry()
	if err != nil {
		return fmt.Errorf("sentry init, %w", err)
	}
	defer sentry.Flush(time.Second * 2)

	start := time.Now()

	port := os.Getenv("PORT")
	if port == "" {
		port = "4400"
	}

	logger := infra.NewLogger()
	if logger == nil {
		return fmt.Errorf("zap:logger is nil")
	}

	tweetColl, err := infra.NewTweetCollection()
	if err != nil {
		return fmt.Errorf("twitter:docstore collection, %w", err)
	}
	defer tweetColl.Close()

	bucket, err := infra.NewBucketFromEnvironment()
	if err != nil {
		return fmt.Errorf("blob bucket, %w", err)
	}

	defer bucket.Close()

	h := handlerImpl{
		log:    logger,
		repo:   tweet.NewRepository(tweetColl),
		bucket: bucket,
	}

	http.HandleFunc("/sitemap", h.handleRequest)

	diff := time.Now().Sub(start)
	logger.Info("initialized objects", Duration("elapsed", diff))

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		return fmt.Errorf("http:ListenAndServe port :%s, %w", port, err)
	}
	return nil
}

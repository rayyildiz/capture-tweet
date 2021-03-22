package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"com.capturetweet/internal/infra"
	"com.capturetweet/pkg/tweet"
	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func init() {
	godotenv.Load()
}

func main() {
	if err := Run(); err != nil {
		log.Fatalf("%v", err)
	}
}

func Run() error {
	infra.RegisterLogger()
	defer sentry.Flush(time.Second * 2)

	start := time.Now()

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
		repo:   tweet.NewRepository(tweetColl),
		bucket: bucket,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/sitemap", h.handleRequest)

	zap.L().Info("initialized objects", zap.Duration("elapsed", time.Since(start).Round(time.Millisecond)))

	port := infra.Port()
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		return fmt.Errorf("http:ListenAndServe port :%s, %w", port, err)
	}
	return nil
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"capturetweet.com/internal/infra"
	"capturetweet.com/pkg/tweet"
	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("can't load .env file, %v", err)
	}
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

	tweetColl := infra.NewTweetCollection()
	defer tweetColl.Close()

	bucket := infra.NewBucketFromEnvironment()
	defer bucket.Close()

	tweetService := tweet.NewService(tweet.NewRepository(tweetColl), nil, nil, nil, nil)
	if tweetService == nil {
		return fmt.Errorf("tweet service is nil")
	}

	h := handlerImpl{
		service: tweetService,
		bucket:  bucket,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/log", h.handleLog)
	mux.HandleFunc("/resize", h.handleResize)

	zap.L().Info("initialized objects", zap.Duration("elapsed", time.Since(start).Round(time.Millisecond)))

	port := infra.Port()
	zap.L().Info("thumb server is starting at port", zap.String("port", port))
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		return fmt.Errorf("http:ListenAndServe port :%s, %w", port, err)
	}
	return nil
}

package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"capturetweet.com/internal/infra"
	"capturetweet.com/pkg/tweet"
	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		slog.Debug("can't load .env file", slog.Any("err", err))
	}
}

func main() {
	if err := Run(); err != nil {
		log.Fatalf("%v", err)
	}
}

func Run() error {
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

	slog.Info("initialized objects", slog.Duration("elapsed", time.Since(start).Round(time.Millisecond)))

	port := infra.Port()
	slog.Info("thumb server is starting at port", slog.String("port", port))
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		return fmt.Errorf("http:ListenAndServe port :%s, %w", port, err)
	}
	return nil
}

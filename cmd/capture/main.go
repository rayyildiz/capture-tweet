package main

import (
	"com.capturetweet/internal/infra"
	"com.capturetweet/pkg/browser"
	"com.capturetweet/pkg/tweet"
	"fmt"
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
	if err := Run(); err != nil {
		log.Printf("%v", err)
		os.Exit(1)
	}
}

func Run() error {
	err := infra.InitSentry()
	if err != nil {
		return fmt.Errorf("sentry init: %w", err)
	}
	defer sentry.Flush(time.Second * 2)

	start := time.Now()
	logger := infra.NewLogger()
	if logger == nil {
		return fmt.Errorf("zap:logger is nil")
	}

	tweetColl, err := infra.NewTweetCollection()
	if err != nil {
		return fmt.Errorf("twitter:docstore collection %w", err)
	}
	defer tweetColl.Close()

	bucket, err := infra.NewBucketFromEnvironment()
	if err != nil {
		return fmt.Errorf("bloc bucket, %w", err)
	}
	defer bucket.Close()

	tweetService := tweet.NewService(tweet.NewRepository(tweetColl), nil, nil, nil, logger, nil)
	if tweetService == nil {
		return fmt.Errorf("tweet service initialize")
	}

	browserService := browser.NewService(logger, tweetService, bucket)
	if browserService == nil {
		return fmt.Errorf("browser service initialize")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "4200"
	}

	h := handlerImpl{
		log:     logger,
		service: browserService,
	}
	http.HandleFunc("/capture", h.handleCapture)

	diff := time.Now().Sub(start)
	logger.Info("initialized objects", zap.Duration("elapsed", diff))

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		return fmt.Errorf("http:ListenAndServe port :%s, %w", port, err)
	}

	return nil
}

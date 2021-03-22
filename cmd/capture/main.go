package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"com.capturetweet/internal/infra"
	"com.capturetweet/pkg/browser"
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
		return fmt.Errorf("twitter:docstore collection %w", err)
	}
	defer tweetColl.Close()

	bucket, err := infra.NewBucketFromEnvironment()
	if err != nil {
		return fmt.Errorf("bloc bucket, %w", err)
	}
	defer bucket.Close()

	tweetService := tweet.NewServiceWithRepository(tweet.NewRepository(tweetColl))
	if tweetService == nil {
		return fmt.Errorf("init tweet service, %w", err)
	}

	browserService := browser.NewService(tweetService, bucket)
	if browserService == nil {
		return fmt.Errorf("browser service initialize")
	}

	h := handlerImpl{
		service: browserService,
	}

	handler := http.NewServeMux()
	handler.HandleFunc("/capture", h.handleCapture)

	zap.L().Info("initialized objects", zap.Duration("elapsed", time.Since(start).Round(time.Millisecond)))

	port := infra.Port()
	err = http.ListenAndServe(":"+port, handler)
	if err != nil {
		return fmt.Errorf("http:ListenAndServe port :%s, %w", port, err)
	}

	return nil
}

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

	tweetColl := infra.NewTweetCollection()
	defer tweetColl.Close()

	bucket := infra.NewBucketFromEnvironment()
	defer bucket.Close()

	tweetService := tweet.NewServiceWithRepository(tweet.NewRepository(tweetColl))
	if tweetService == nil {
		return fmt.Errorf("init tweet service")
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
	zap.L().Info("capture server is starting at port", zap.String("port", port))

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		return fmt.Errorf("http:ListenAndServe port :%s, %w", port, err)
	}

	return nil
}

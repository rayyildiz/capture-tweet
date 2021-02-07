package main

import (
	"com.capturetweet/internal/infra"
	"com.capturetweet/pkg/browser"
	"com.capturetweet/pkg/tweet"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/run"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
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

	err := infra.InitSentry()
	if err != nil {
		return fmt.Errorf("sentry init: %w", err)
	}
	defer sentry.Flush(time.Second * 2)

	telemetryClose := infra.NewTelemetry()
	defer telemetryClose()

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
		tracer:  otel.Tracer("cmd/capture"),
	}
	http.HandleFunc("/capture", h.handleCapture)

	zap.L().Info("initialized objects", zap.Duration("elapsed", time.Since(start)))

	port := run.Port()
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		return fmt.Errorf("http:ListenAndServe port :%s, %w", port, err)
	}

	return nil
}

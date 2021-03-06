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
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
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

	telemetryClose := infra.NewTelemetry()
	defer telemetryClose()

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

	tweetService := tweet.NewService(tweet.NewRepository(tweetColl), nil, nil, nil, nil)
	if tweetService == nil {
		return fmt.Errorf("tweet service is nil")
	}

	h := handlerImpl{
		service: tweetService,
		bucket:  bucket,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/resize", h.handleResize)

	zap.L().Info("initialized objects", zap.Duration("elapsed", time.Since(start)))

	port := infra.Port()
	err = http.ListenAndServe(":"+port, otelhttp.NewHandler(mux, "thumb"))
	if err != nil {
		return fmt.Errorf("http:ListenAndServe port :%s, %w", port, err)
	}
	return nil
}

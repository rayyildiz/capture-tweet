package main

import (
	"com.capturetweet/internal/infra"
	"com.capturetweet/pkg/tweet"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/run"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
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
		return fmt.Errorf("sentry init, %w", err)
	}
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

	h := handlerImpl{
		repo:   tweet.NewRepository(tweetColl),
		bucket: bucket,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/sitemap", h.handleRequest)

	zap.L().Info("initialized objects", zap.Duration("elapsed", time.Since(start)))

	port := run.Port()
	err = http.ListenAndServe(":"+port, otelhttp.NewHandler(mux, "sitemap"))
	if err != nil {
		return fmt.Errorf("http:ListenAndServe port :%s, %w", port, err)
	}
	return nil
}

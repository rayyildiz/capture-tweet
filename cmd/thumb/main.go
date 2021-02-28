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
	"runtime/debug"
	"time"
)

var version string // do not remove or modify

func init() {
	godotenv.Load()
}

func init() {
	if version == "" {
		i, ok := debug.ReadBuildInfo()
		if !ok {
			return
		}
		version = i.Main.Version
	}
}

func main() {
	if err := Run(); err != nil {
		log.Fatalf("%v", err)
	}
}

func Run() error {
	infra.RegisterLogger(version)

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

	port := run.Port()
	err = http.ListenAndServe(":"+port, otelhttp.NewHandler(mux, "thumb"))
	if err != nil {
		return fmt.Errorf("http:ListenAndServe port :%s, %w", port, err)
	}
	return nil
}

package main

import (
	"com.capturetweet/internal/infra"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/run"
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

	telemetryClose := infra.NewTelemetry()
	defer telemetryClose()

	start := time.Now()

	bucket, err := infra.NewBucketFromEnvironment()
	if err != nil {
		return fmt.Errorf("blob bucket, %w", err)
	}
	defer bucket.Close()

	h := handlerImpl{
		bucket: bucket,
	}
	http.HandleFunc("/dlq-capture", h.handleMessages)

	zap.L().Info("initialized objects", zap.Duration("elapsed", time.Since(start)))

	port := run.Port()
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		return fmt.Errorf("http:ListenAndServe port :%s, %w", port, err)
	}
	return nil
}

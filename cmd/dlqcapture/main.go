package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"com.capturetweet/internal/infra"
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

	mux := http.NewServeMux()

	bucket, err := infra.NewBucketFromEnvironment()
	if err != nil {
		return fmt.Errorf("blob bucket, %w", err)
	}
	defer bucket.Close()

	h := handlerImpl{
		bucket: bucket,
	}
	mux.HandleFunc("/dlq-capture", h.handleMessages)

	zap.L().Info("initialized objects", zap.Duration("elapsed", time.Since(start).Round(time.Millisecond)))

	port := infra.Port()
	zap.L().Info("dlq-capture server is starting at port", zap.String("port", port))

	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		return fmt.Errorf("http:ListenAndServe port :%s, %w", port, err)
	}
	return nil
}

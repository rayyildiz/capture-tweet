package main

import (
	"com.capturetweet/internal/infra"
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
		log.Fatalf("%v", err)
	}
}

func Run() error {
	err := infra.InitSentry()
	if err != nil {
		return fmt.Errorf("sentry init, %w", err)
	}
	defer sentry.Flush(time.Second * 2)

	start := time.Now()
	infra.RegisterLogger()

	bucket, err := infra.NewBucketFromEnvironment()
	if err != nil {
		return fmt.Errorf("blob bucket, %w", err)
	}
	defer bucket.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "4400"
	}

	h := handlerImpl{
		bucket: bucket,
	}
	http.HandleFunc("/dlq-capture", h.handleMessages)

	diff := time.Now().Sub(start)
	zap.L().Info("initialized objects", zap.Duration("elapsed", diff))

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		return fmt.Errorf("http:ListenAndServe port :%s, %w", port, err)
	}
	return nil
}

package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	"capturetweet.com/internal/infra"
	"github.com/getsentry/sentry-go"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		slog.Debug("can't load .env file", slog.Any("err", err))
	}
}

func main() {
	if err := Run(); err != nil {
		log.Fatalf("%v", err)
	}
}

func Run() error {
	defer sentry.Flush(time.Second * 2)

	start := time.Now()

	h := handlerImpl{
		service: initializeBrowserService(),
	}

	handler := http.NewServeMux()
	handler.HandleFunc("ANY /capture", h.handleCapture)

	slog.Info("initialized objects", slog.Duration("elapsed", time.Since(start).Round(time.Millisecond)))

	port := infra.Port()
	slog.Info("capture server is starting at port", slog.String("port", port))
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		return fmt.Errorf("http:ListenAndServe port :%s, %w", port, err)
	}

	return nil
}

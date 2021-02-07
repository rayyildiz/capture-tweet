package infra

import (
	"github.com/getsentry/sentry-go"
	"github.com/kelseyhightower/run"
	"os"
)

func InitSentry() error {
	dsn := os.Getenv("SENTRY_DSN")
	if len(dsn) > 0 {
		return sentry.Init(sentry.ClientOptions{
			Dsn:     dsn,
			Dist:    run.ServiceName(),
			Release: run.Revision(),
		})
	}
	return nil
}

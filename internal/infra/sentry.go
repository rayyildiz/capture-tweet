package infra

import (
	"github.com/getsentry/sentry-go"
	"github.com/kelseyhightower/run"
	"os"
)

func InitSentry(version string) error {
	if len(version) < 2 {
		version = run.Revision()
	}
	dsn := os.Getenv("SENTRY_DSN")
	if len(dsn) > 0 {
		return sentry.Init(sentry.ClientOptions{
			Dsn:     dsn,
			Dist:    run.ServiceName(),
			Release: version,
		})
	}
	return nil
}

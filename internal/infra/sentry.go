package infra

import (
	"os"

	"github.com/getsentry/sentry-go"
)

func InitSentry() error {
	version := Version
	if len(version) < 2 {
		version = Revision()
	}
	dsn := os.Getenv("SENTRY_DSN")
	if len(dsn) > 0 {
		return sentry.Init(sentry.ClientOptions{
			Dsn:     dsn,
			Dist:    ServiceName(),
			Release: version,
		})
	}
	return nil
}

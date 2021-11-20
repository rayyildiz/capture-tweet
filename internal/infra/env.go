package infra

import "os"

// Port returns the port your HTTP server should listen on. Default port 4000.
func Port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	return port
}

// Revision returns the name of the Cloud Run revision being run.
func Revision() string {
	return os.Getenv("K_REVISION")
}

// Configuration returns the name of the Cloud Run configuration being run.
func Configuration() string {
	return os.Getenv("K_CONFIGURATION")
}

// ServiceName returns the name of the Cloud Run service being run.
func ServiceName() string {
	return os.Getenv("K_SERVICE")
}

// ProjectID returns teh project id of GCP.
func ProjectID() string {
	return os.Getenv("GOOGLE_CLOUD_PROJECT")
}

// SentryRelease is for sentry
var sentryRelease = "0.5.0"

func SentryRelease() string {
	v := os.Getenv("SENTRY_VERSION")
	if len(v) > 1 {
		v = sentryRelease
	}
	return v
}

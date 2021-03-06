package infra

import "os"

// Port returns the port your HTTP server should listen on.
func Port() string {
	return os.Getenv("PORT")
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

func SentryDSN() string {
	return os.Getenv("SENTRY_DSN")
}

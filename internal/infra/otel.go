package infra

import (
	"context"
	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"github.com/getsentry/sentry-go"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
	"log"
)

func NewTelemetry() func() {
	projectId := ProjectID()
	if len(projectId) > 0 {
		zap.L().Info("injection google telemetry telemetry", zap.String("projectId", projectId))
		exporter, err := texporter.NewExporter(texporter.WithProjectID(projectId))
		if err != nil {
			sentry.CaptureException(err)
			log.Fatalf("could not create exportter")
		}

		tp := sdktrace.NewTracerProvider(sdktrace.WithSyncer(exporter))
		otel.SetTracerProvider(tp)

		return func() {
			exporter.Shutdown(context.Background())
		}
	}
	return func() {}
}

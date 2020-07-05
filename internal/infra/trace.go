package infra

import (
	"fmt"
	export "go.opentelemetry.io/otel/sdk/export/trace"
	"os"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/exporters/trace/stdout"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func InitTrace() error {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

	var exporter export.SpanSyncer
	if len(projectID) > 0 {
		exo, err := texporter.NewExporter(texporter.WithProjectID(projectID))
		if err != nil {
			return fmt.Errorf("texporter.NewExporter: %v", err)
		}
		exporter = exo

	} else {
		exo, err := stdout.NewExporter(stdout.Options{PrettyPrint: true})
		if err != nil {
			return err
		}

		exporter = exo
	}

	tp, err := sdktrace.NewProvider(sdktrace.WithSyncer(exporter))
	if err != nil {
		return err
	}
	global.SetTraceProvider(tp)

	return nil
}

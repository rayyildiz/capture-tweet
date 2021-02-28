package infra

import (
	"github.com/kelseyhightower/run"
	"github.com/lightstep/otel-launcher-go/launcher"
	"go.uber.org/zap"
	"os"
)

func NewTelemetry() func() {

	token := os.Getenv("LIGHSTEP_TOKEN")
	if len(token) > 0 {
		zap.L().Info("injection lighstep telemetry")
		cfg := launcher.ConfigureOpentelemetry(
			launcher.WithServiceName(run.ServiceName()),
			launcher.WithServiceVersion(run.Revision()),
			launcher.WithAccessToken(token),
		)
		return func() {
			cfg.Shutdown()
		}
	}
	return func() {}
}

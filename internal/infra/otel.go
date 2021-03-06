package infra

import (
	"os"

	"github.com/lightstep/otel-launcher-go/launcher"
	"go.uber.org/zap"
)

func NewTelemetry() func() {

	token := os.Getenv("LIGHSTEP_TOKEN")
	if len(token) > 0 {
		zap.L().Info("injection lighstep telemetry")
		cfg := launcher.ConfigureOpentelemetry(
			launcher.WithServiceName(ServiceName()),
			launcher.WithServiceVersion(Revision()),
			launcher.WithAccessToken(token),
		)
		return func() {
			cfg.Shutdown()
		}
	}
	return func() {}
}

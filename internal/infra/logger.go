package infra

import (
	"fmt"
	"github.com/kelseyhightower/run"
	"go.uber.org/zap"
)

func RegisterLogger(version string) {
	var config zap.Config

	if IsDebug() {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	config.InitialFields = map[string]interface{}{
		"version":      version,
		"service_name": run.ServiceName(),
	}
	logger, err := config.Build()
	if err != nil {
		fmt.Printf("Could not create zap logger, %v\n", err)
		return
	}

	zap.ReplaceGlobals(logger)
}

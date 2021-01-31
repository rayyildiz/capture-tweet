package infra

import (
	"fmt"
	"go.uber.org/zap"
)

func RegisterLogger() {
	var config zap.Config

	if IsDebug() {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	logger, err := config.Build()
	if err != nil {
		fmt.Printf("Could not create zap logger, %v\n", err)
		return
	}
	zap.ReplaceGlobals(logger)
}

package infra

import (
	"fmt"
	"go.uber.org/zap/zapcore"

	"github.com/TheZeroSlave/zapsentry"
	"go.uber.org/zap"
)

func RegisterLogger() {
	var config zap.Config

	if IsDebug() {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}

	config.InitialFields = map[string]interface{}{
		"version":     Version,
		"serviceName": ServiceName(),
	}
	logger, err := config.Build()
	if err != nil {
		fmt.Printf("Could not create zap logger, %v\n", err)
		return
	}
	logger = modifyToSentryLogger(logger)

	zap.ReplaceGlobals(logger)
}

func modifyToSentryLogger(log *zap.Logger) *zap.Logger {
	dsn := SentryDSN()
	if len(dsn) < 2 {
		return log
	}

	cfg := zapsentry.Configuration{
		Level: zapcore.ErrorLevel, //when to send message to sentry
		Tags: map[string]string{
			"serviceName": ServiceName(),
			"revision":    Revision(),
			"version":     Version,
		},
	}
	core, err := zapsentry.NewCore(cfg, zapsentry.NewSentryClientFromDSN(dsn))
	if err != nil {
		log.Warn("failed to init zap", zap.Error(err))
	}
	return zapsentry.AttachCoreToLogger(core, log)
}

package infra

import (
	"fmt"
	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	logger, err := zap.NewProductionConfig().Build()
	if err != nil {
		fmt.Printf("Coudl not create zap logger, %v\n", err)
		return nil
	}
	return logger
}

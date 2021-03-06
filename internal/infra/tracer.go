package infra

import (
	"context"
	"go.uber.org/zap"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type ZapLogger struct{}

var _ interface {
	graphql.HandlerExtension
	graphql.ResponseInterceptor
} = ZapLogger{}

func (a ZapLogger) ExtensionName() string {
	return "ZapLogger"
}

func (a ZapLogger) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

func (a ZapLogger) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	start := time.Now()
	oc := graphql.GetOperationContext(ctx)
	response := next(ctx)
	if oc.OperationName != "IntrospectionQuery" && oc.OperationName != "WatchChange" {
		if len(response.Errors) > 0 {
			zap.L().Warn("request:error", zap.Bool("is_error", true), zap.Duration("elapsed", time.Since(start).Round(time.Millisecond)), zap.String("operation_name", oc.OperationName), zap.String("error", response.Errors.Error()))
		} else {
			zap.L().Info("request:success", zap.Bool("is_error", false), zap.Duration("elapsed", time.Since(start).Round(time.Millisecond)), zap.String("operation_name", oc.OperationName))
		}
	}
	return response
}

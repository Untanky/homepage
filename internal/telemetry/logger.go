package telemetry

import (
	"context"
	"log/slog"
)

type contextKey int

const contextTraceID contextKey = iota

type Config struct {
	LogLevel string
}

type contextHandler struct {
	slog.Handler
}

func (handler contextHandler) Handle(ctx context.Context, record slog.Record) error {
	if ctx == nil {
		return handler.Handler.Handle(ctx, record)
	}

	if traceID, ok := ctx.Value(contextTraceID).(string); ok {
		record.AddAttrs(slog.String("traceID", traceID))
	}

	return handler.Handler.Handle(ctx, record)
}

func Setup(ctx context.Context, config Config) (*slog.Logger, error) {
	handler := slog.Default().Handler()

	logger := slog.New(contextHandler{handler})

	return logger, nil
}

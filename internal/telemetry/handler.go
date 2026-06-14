package telemetry

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"log/slog"
	"math/rand/v2"
	"net/http"
)

type Handler struct {
	next http.Handler

	logger *slog.Logger
	rand   *rand.Rand
}

func NewHandler(next http.Handler, logger *slog.Logger) Handler {
	return Handler{
		next:   next,
		logger: logger,
		rand:   rand.New(rand.NewPCG(0, 1)),
	}
}

func (handler Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	traceID := handler.generateTraceID()

	ctx := context.WithValue(request.Context(), contextTraceID, traceID)
	request.WithContext(ctx)
	writer.Header().Set("x-trace-id", traceID)

	handler.logger.InfoContext(ctx, "received request",
		slog.String("method", request.Method),
		slog.String("path", request.URL.Path),
	)

	handler.next.ServeHTTP(writer, request)
}

func (handler Handler) generateTraceID() string {
	traceBytes := make([]byte, 0, 24)
	buffer := make([]byte, 0, 8)

	order := binary.NativeEndian

	traceBytes = hex.AppendEncode(traceBytes, order.AppendUint64(buffer[:0], handler.rand.Uint64()))
	traceBytes = hex.AppendEncode(traceBytes, order.AppendUint64(buffer[:0], handler.rand.Uint64()))

	return string(traceBytes)
}

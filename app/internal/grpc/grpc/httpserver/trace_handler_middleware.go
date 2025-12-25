package httpserver

import (
	"context"
	"net/http"

	servicetrace "github.com/dzamyatin/atomWebsite/internal/service/trace"
)

type TraceHandlerMiddleware struct {
	h     http.Handler
	trace *servicetrace.Trace
}

func NewTraceHandlerMiddleware(
	h http.Handler,
	trace *servicetrace.Trace,
) *TraceHandlerMiddleware {
	return &TraceHandlerMiddleware{h: h, trace: trace}
}

func (h *TraceHandlerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.trace.Trace(
		r.Context(),
		"http",
		func(ctx context.Context) {
			h.h.ServeHTTP(w, r)
		},
	)
}

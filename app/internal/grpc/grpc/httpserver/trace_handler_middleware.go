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
	var tags []servicetrace.Tag

	for headerName, headerValues := range r.Header {
		for _, headerValue := range headerValues {
			tags = append(tags, servicetrace.Tag{
				Key:   headerName,
				Value: headerValue,
			})
		}
	}

	tags = append(
		tags,
		[]servicetrace.Tag{
			{
				Key:   "method",
				Value: r.Method,
			},
			{
				Key:   "url",
				Value: r.RequestURI,
			},
		}...,
	)

	h.trace.Trace(
		r.Context(),
		"http",
		func(ctx context.Context) {
			h.h.ServeHTTP(w, r.WithContext(ctx))
		},
		tags...,
	)
}

package httpserver

import (
	"net/http"

	"github.com/dzamyatin/atomWebsite/internal/service/metric"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type MetricHandlerMiddleware struct {
	h      http.Handler
	metric *metric.Metric
}

func NewMetricHandlerMiddleware(h http.Handler, metric *metric.Metric) *MetricHandlerMiddleware {
	return &MetricHandlerMiddleware{h: h, metric: metric}
}

func (h *MetricHandlerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.metric.IncomingRequestHistogram(
		func() {
			h.h.ServeHTTP(w, r)
		},
	)
}

type MetricMuxHandlerMiddleware struct {
	metric *metric.Metric
}

func NewMetricMuxHandlerMiddleware(metric *metric.Metric) *MetricMuxHandlerMiddleware {
	return &MetricMuxHandlerMiddleware{metric: metric}
}

func (r *MetricMuxHandlerMiddleware) Middleware() func(h runtime.HandlerFunc) runtime.HandlerFunc {
	return func(h runtime.HandlerFunc) runtime.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
			r.metric.IncomingRequestHistogram(
				func() {
					h(w, req, pathParams)
					//
					//ctx := req.Context()
					//
					//if p, ok := runtime.HTTPPattern(ctx); ok {
					//	_ = p
					//}
					//_ = ctx
				},
			)
		}
	}
}

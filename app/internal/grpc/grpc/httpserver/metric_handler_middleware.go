package httpserver

import (
	"net/http"

	"github.com/dzamyatin/atomWebsite/internal/service/metric"
)

type MetricHandlerMiddleware struct {
	h      http.Handler
	metric *metric.Metric
}

func NewMetricHandlerMiddleware(h http.Handler, metric *metric.Metric) *MetricHandlerMiddleware {
	return &MetricHandlerMiddleware{h: h, metric: metric}
}

func (h *MetricHandlerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("access-control-allow-credentials", "true")
	w.Header().Set("access-control-allow-headers", "Content-Type, Authorization")
	w.Header().Set("access-control-allow-methods", "PUT, GET, POST, PATCH, DELETE, OPTIONS")
	w.Header().Set("access-control-allow-origin", "http://localhost:5173")
	w.Header().Set("access-control-max-age", "1728000")

	h.metric.IncomingRequestHistogram(func() {
		h.h.ServeHTTP(w, r)
	})
}

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
	h.metric.IncomingRequestHistogram(func() {
		h.h.ServeHTTP(w, r)
	})
}

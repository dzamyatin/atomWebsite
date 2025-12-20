package metric

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	_ "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

const promTimeout = 10 * time.Second

var defaultBuckets = []float64{.01, .02, .05, .1, .25, .5, 1, 2.5, 5, 10}

type Registry struct {
	registry *prometheus.Registry
	logger   *zap.Logger
}

func NewRegistry(logger *zap.Logger) *Registry {
	return &Registry{
		registry: prometheus.NewRegistry(),
		logger:   logger,
	}
}

func (r *Registry) RunGCMetrics() error {
	return r.registry.Register(collectors.NewGoCollector())
}

func (r *Registry) HTTPHandler() http.Handler {
	return promhttp.InstrumentMetricHandler(
		r.registry, promhttp.HandlerFor(
			r.registry,
			promhttp.HandlerOpts{
				Timeout: promTimeout,
			},
		),
	)
}

type MeasuredFunc func()
type MetricFunc func(measured MeasuredFunc)

func (r *Registry) Histogram(
	name string,
	labels prometheus.Labels,
) MetricFunc {
	hist := prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:        name,
			ConstLabels: labels,
			Buckets:     defaultBuckets,
		},
	)
	err := r.registry.Register(hist)

	if err != nil {
		r.logger.Error("failed to register histogram", zap.Error(err))
	}

	return func(f MeasuredFunc) {
		b := time.Now()
		f()
		hist.Observe(float64(time.Since(b).Nanoseconds() / 1000000000))
	}
}

package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	_ "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"net/http"
	"time"
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

func (r *Registry) Histogram(
	f MeasuredFunc,
	name string,
	labels prometheus.Labels,
) {
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

	f()

	b := time.Now()

	hist.Observe(float64(time.Since(b).Nanoseconds()))
}

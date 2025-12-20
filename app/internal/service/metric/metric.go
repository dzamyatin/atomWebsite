package metric

import (
	"go.uber.org/zap"
)

const (
	nameIncomingGRPCRequestHistogram = "incoming_grpc_request_latency"
)

type Metric struct {
	logger *zap.Logger

	incomingGRPCRequestHistogram MetricFunc
}

func NewMetric(
	logger *zap.Logger,
	registry *Registry,
) *Metric {
	m := &Metric{
		logger: logger,
	}

	m.incomingGRPCRequestHistogram = registry.Histogram(
		nameIncomingGRPCRequestHistogram,
		nil,
	)

	return m
}

func (r *Metric) IncomingRequestHistogram(f MeasuredFunc) {
	r.incomingGRPCRequestHistogram(f)
}

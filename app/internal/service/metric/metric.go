package metric

import (
	"go.uber.org/zap"
)

const (
	nameIncomingGRPCRequestHistogram = "incoming_grpc_request_latency"
)

type Metric struct {
	logger   *zap.Logger
	registry *Registry
}

func NewMetric(
	logger *zap.Logger,
	registry *Registry,
) *Metric {
	return &Metric{
		logger:   logger,
		registry: registry,
	}
}

func (m *Metric) IncomingRequestHistogram(f MeasuredFunc) {
	m.registry.Histogram(
		f,
		nameIncomingGRPCRequestHistogram,
		nil,
	)
}

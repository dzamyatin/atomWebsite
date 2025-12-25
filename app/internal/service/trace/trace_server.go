package servicetrace

import (
	"context"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/uber/jaeger-client-go/config"
	zap2 "github.com/uber/jaeger-client-go/log/zap"
	"go.uber.org/zap"
)

const (
	serviceName = "website"
)

type TraceServer struct {
	logger *zap.Logger
	tracer opentracing.Tracer
	cfg    config.Configuration
	closer io.Closer
	ch     chan struct{}
}

func NewTraceServer(
	logger *zap.Logger,
	agentHost string,
) *TraceServer {
	traceServer := &TraceServer{
		ch:     make(chan struct{}, 1),
		tracer: opentracing.NoopTracer{},
		logger: logger,
		cfg: config.Configuration{
			ServiceName: serviceName,
			Sampler: &config.SamplerConfig{
				Type:  "const",
				Param: 1,
			},
			Reporter: &config.ReporterConfig{
				LogSpans:           true,
				LocalAgentHostPort: agentHost,
				//CollectorEndpoint:  "",
			},
		},
	}

	return traceServer
}

func (r *TraceServer) GetTrace() *opentracing.Tracer {
	return &r.tracer
}

func (r *TraceServer) Run() error {
	//tracer, closer, err := r.cfg.NewTracer(config.Logger(jaeger.StdLogger))
	tracer, closer, err := r.cfg.NewTracer(
		config.Logger(
			zap2.NewLogger(r.logger),
		),
	)
	if err != nil {
		return errors.Wrap(err, "start tracer")
	}

	r.closer = closer
	r.tracer = tracer

	return nil
}

func (r *TraceServer) Start(ctx context.Context) error {
	select {
	case <-r.ch:
		r.logger.Info("Tracing server shutting down by shutting down func")
		return nil
	case <-ctx.Done():
		r.logger.Info("Tracing server shutting down by canceling context")
		return ctx.Err()
	}
}

func (r *TraceServer) Shutdown(ctx context.Context) error {
	r.ch <- struct{}{}
	if r.closer == nil {
		return nil
	}

	return r.closer.Close()
}

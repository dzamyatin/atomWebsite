package executors

import (
	"context"
	"net"
	"net/http"

	grpcservice2 "github.com/dzamyatin/atomWebsite/internal/grpc/grpc"
	mainarg "github.com/dzamyatin/atomWebsite/internal/service/arg"
	"github.com/dzamyatin/atomWebsite/internal/service/config"
	"github.com/dzamyatin/atomWebsite/internal/service/metric"
	"github.com/dzamyatin/atomWebsite/internal/service/process"
	servicetrace "github.com/dzamyatin/atomWebsite/internal/service/trace"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ArgGrpcProcess struct {
	mainarg.Arg
}

type GrpcProcessCommand struct {
	logger         *zap.Logger
	processManager *process.ProcessShutdownerManager
	serv           *grpcservice2.GRPCServer
	listener       *process.SignalListener
	metric         *metric.Registry
	server         *grpcservice2.HTTPServer
	config         config.AppConfig
	trace          *servicetrace.TraceServer
}

func NewGrpcProcessCommand(
	logger *zap.Logger,
	processManager *process.ProcessShutdownerManager,
	serv *grpcservice2.GRPCServer,
	listener *process.SignalListener,
	metric *metric.Registry,
	server *grpcservice2.HTTPServer,
	cfg config.AppConfig,
	trace *servicetrace.TraceServer,
) *GrpcProcessCommand {
	return &GrpcProcessCommand{
		logger:         logger,
		processManager: processManager,
		serv:           serv,
		listener:       listener,
		metric:         metric,
		server:         server,
		config:         cfg,
		trace:          trace,
	}
}

func (r *GrpcProcessCommand) Execute(ctx context.Context, u ArgGrpcProcess) error {
	return r.processManager.Run(
		ctx,
		process.NewProcessIniter(
			"tracer",
			func(ctx context.Context) (process.ProcessStarter, error) {
				if err := r.trace.Run(); err != nil {
					return nil, errors.Wrap(err, "tracer error")
				}

				return r.trace, nil
			},
		),
		process.NewProcess("grpc server", r.serv),
		process.NewProcess("http server", r.server),
		process.NewProcess("signal listener", r.listener),
		process.NewProcess("metric server", r.newHttpMetricProcessor(r.metric)),
	)
}

func (r *GrpcProcessCommand) newHttpMetricProcessor(registry *metric.Registry) *process.Processor {
	mux := http.NewServeMux()
	mux.Handle("/metrics", registry.HTTPHandler())

	ctx := context.Background()

	server := &http.Server{
		Addr:    r.config.AddrMetric,
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	return process.NewProcessor(
		func(_ context.Context) error {
			return server.ListenAndServe()
		},
		func() error {
			return server.Shutdown(ctx)
		},
	)
}

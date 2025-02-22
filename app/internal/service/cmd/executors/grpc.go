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
}

func NewGrpcProcessCommand(
	logger *zap.Logger,
	processManager *process.ProcessShutdownerManager,
	serv *grpcservice2.GRPCServer,
	listener *process.SignalListener,
	metric *metric.Registry,
	server *grpcservice2.HTTPServer,
	cfg config.AppConfig,
) *GrpcProcessCommand {
	return &GrpcProcessCommand{
		logger:         logger,
		processManager: processManager,
		serv:           serv,
		listener:       listener,
		metric:         metric,
		server:         server,
		config:         cfg,
	}
}

func (r *GrpcProcessCommand) Execute(ctx context.Context, u ArgGrpcProcess) error {
	return r.processManager.Run(
		ctx,
		process.Process{
			Name:   "grpc server",
			Object: r.serv,
		},
		process.Process{
			Name:   "http server",
			Object: r.server,
		},
		process.Process{
			Name:   "signal listener",
			Object: r.listener,
		},
		process.Process{
			Name:   "http listener",
			Object: r.newHttpMetricProcessor(r.metric),
		},
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

package di

import (
	"context"
	"database/sql"
	atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"
	grpcservice2 "github.com/dzamyatin/atomWebsite/internal/grpc/grpc"
	"github.com/dzamyatin/atomWebsite/internal/service/metric"
	"github.com/dzamyatin/atomWebsite/internal/service/process"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

func newGRPCProcessManager(
	logger *zap.Logger,
	serv *grpcservice2.GRPCServer,
	listener *process.SignalListener,
	db *sql.DB,
	metric *metric.Registry,
) *process.ProcessManager {
	return process.NewProcessManager(
		logger,
		process.Process{
			Name:   "grpc server",
			Object: serv,
		},
		process.Process{
			Name:   "signal listener",
			Object: listener,
		},
		process.Process{
			Name: "database shutdowner",
			Object: process.NewShutdowner(
				logger,
				func() error {
					return db.Close()
				},
			),
		},
		process.Process{
			Name:   "http listener",
			Object: newHttpMetricProcessor(metric),
		},
	)
}

func newDb() (*sql.DB, error) {
	config := getConfig().Db

	query := "dbname=" + config.Database
	if !config.SSL {
		query += "&sslmode=disable"
	}

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(config.Username, config.Password),
		Host:     config.Host + ":" + config.Port,
		RawQuery: query,
	}

	db, err := sql.Open("postgres", u.String())

	if err != nil {
		return nil, errors.Wrap(err, "could not connect to postgres")
	}

	err = db.Ping()

	if err != nil {
		return nil, errors.Wrap(err, "could not connect to postgres")
	}

	return db, nil
}

func newLogger() *zap.Logger {
	return zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.AddSync(os.Stdout),
			zapcore.InfoLevel,
		),
	)
}

func newServer(
	grpcServer *grpc.Server,
) *grpcservice2.GRPCServer {
	lis, err := net.Listen("tcp", getConfig().AddrGrpc)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return grpcservice2.NewGRPCServer(
		lis,
		grpcServer,
	)
}

func newGrpcServer(
	server grpcservice2.AuthServer,
	m *metric.Metric,
) *grpc.Server {
	grpcServer := grpc.NewServer()

	grpc.WithIdleTimeout(time.Minute * 1)
	grpc.WithTimeout(time.Minute * 1)

	grpc.WithChainUnaryInterceptor(
		func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
			var err error

			m.IncomingRequestHistogram(
				func() {
					err = invoker(ctx, method, req, reply, cc, opts...)
				},
			)

			return err
		},
	)

	atomWebsite.RegisterAuthServer(grpcServer, server)

	return grpcServer
}

func newHttpMetricProcessor(r *metric.Registry) *process.Processor {
	mux := http.NewServeMux()
	mux.Handle("/metrics", r.HTTPHandler())

	ctx := context.Background()

	server := &http.Server{
		Addr:    getConfig().AddrMetric,
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

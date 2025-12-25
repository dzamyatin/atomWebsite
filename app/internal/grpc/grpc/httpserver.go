package grpc

import (
	"context"
	"net"
	"net/http"
	"time"

	"github.com/dzamyatin/atomWebsite/internal/grpc/grpc/httpserver"
	"github.com/dzamyatin/atomWebsite/internal/service/metric"
	servicetrace "github.com/dzamyatin/atomWebsite/internal/service/trace"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Option func(*HTTPServer)

func WithTimeout(
	readTimeout,
	writeTimeout,
	idleTimeout time.Duration,
) Option {
	return func(s *HTTPServer) {
		s.readTimeout = readTimeout
		s.writeTimeout = writeTimeout
		s.idleTimeout = idleTimeout
	}
}

func WithCors(
	host string,
) Option {
	return func(s *HTTPServer) {
		s.corsHost = host
	}
}

type HTTPServer struct {
	logger       *zap.Logger
	server       *http.Server
	httpAddr     string
	router       *HttpRouter
	metric       *metric.Metric
	readTimeout  time.Duration
	writeTimeout time.Duration
	idleTimeout  time.Duration
	trace        *servicetrace.Trace
	corsHost     string
}

func NewHTTPServer(
	logger *zap.Logger,
	httpAddr string,
	router *HttpRouter,
	metric *metric.Metric,
	trace *servicetrace.Trace,
	options ...Option,
) *HTTPServer {
	s := &HTTPServer{
		logger:   logger,
		httpAddr: httpAddr,
		router:   router,
		metric:   metric,
		trace:    trace,
	}

	for _, option := range options {
		option(s)
	}

	return s
}

func (r *HTTPServer) Shutdown() error {
	if r.server == nil {
		return nil
	}

	return r.server.Shutdown(context.Background())
}

func (r *HTTPServer) Start(ctx context.Context) error {
	r.logger.Info("Starting HTTP server", zap.String("http_addr", r.httpAddr))

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	err := r.router.Apply(ctx, mux)
	if err != nil {
		return errors.Wrap(err, "failed to start HTTP server")
	}

	r.server = &http.Server{
		ReadTimeout:  r.readTimeout,
		WriteTimeout: r.writeTimeout,
		IdleTimeout:  r.idleTimeout,
		Addr:         r.httpAddr,
		Handler: httpserver.NewTraceHandlerMiddleware(
			httpserver.NewMetricHandlerMiddleware(
				httpserver.NewCorsHandlerMiddleware(
					mux,
					r.corsHost,
				),
				r.metric,
			),
			r.trace,
		),
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	return r.server.ListenAndServe()
}

package grpc

import (
	"context"
	"net"
	"net/http"

	"github.com/dzamyatin/atomWebsite/internal/service/metric"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type HTTPServer struct {
	logger   *zap.Logger
	server   *http.Server
	httpAddr string
	router   *HttpRouter
	metric   *metric.Metric
}

func NewHTTPServer(
	logger *zap.Logger,
	httpAddr string,
	router *HttpRouter,
	metric *metric.Metric,
) *HTTPServer {
	return &HTTPServer{
		logger:   logger,
		httpAddr: httpAddr,
		router:   router,
		metric:   metric,
	}
}

func (r *HTTPServer) Shutdown() error {
	if r.server == nil {
		return nil
	}

	return r.server.Shutdown(context.Background())
}

type Handler struct {
	h      http.Handler
	metric *metric.Metric
}

func NewHandler(h http.Handler, metric *metric.Metric) *Handler {
	return &Handler{h: h, metric: metric}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("access-control-allow-credentials", "true")
	w.Header().Set("access-control-allow-headers", "Content-Type, Authorization")
	w.Header().Set("access-control-allow-methods", "PUT, GET, POST, PATCH, DELETE, OPTIONS")
	w.Header().Set("access-control-allow-origin", "http://localhost:5173")
	w.Header().Set("access-control-max-age", "1728000")

	h.metric.IncomingRequestHistogram(func() {
		h.h.ServeHTTP(w, r)
	})
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
		//ReadTimeout:  5 * time.Second,
		//WriteTimeout: 5 * time.Second,
		//IdleTimeout:  5 * time.Second,
		Addr:    r.httpAddr,
		Handler: NewHandler(mux, r.metric),
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	return r.server.ListenAndServe()
}

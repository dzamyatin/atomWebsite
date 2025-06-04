package grpc

import (
	"context"
	atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net"
	"net/http"
)

type HTTPServer struct {
	service  AuthServer
	server   *http.Server
	httpAddr string
}

func NewHTTPServer(
	server AuthServer,
	httpAddr string,
) *HTTPServer {
	return &HTTPServer{
		service:  server,
		httpAddr: httpAddr,
	}
}

func (r *HTTPServer) Shutdown() error {
	if r.server == nil {
		return nil
	}

	return r.server.Shutdown(context.Background())
}

func (r *HTTPServer) Start(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	err := atomWebsite.RegisterAuthHandlerServer(ctx, mux, r.service)
	if err != nil {
		return err
	}

	r.server = &http.Server{
		Addr:    r.httpAddr,
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	return r.server.ListenAndServe()
}

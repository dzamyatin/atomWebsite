package grpc

import (
	"context"
	atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"
	"github.com/dzamyatin/atomWebsite/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
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

type Handler struct {
	h http.Handler
}

func NewHandler(h http.Handler) *Handler {
	return &Handler{h: h}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("access-control-allow-credentials", "true")
	w.Header().Set("access-control-allow-headers", "Content-Type, Authorization")
	w.Header().Set("access-control-allow-methods", "PUT, GET, POST, PATCH, DELETE, OPTIONS")
	w.Header().Set("access-control-allow-origin", "http://localhost:5173")
	w.Header().Set("access-control-max-age", "1728000")

	h.h.ServeHTTP(w, r)
}

func (r *HTTPServer) Start(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	err := atomWebsite.RegisterAuthHandlerServer(ctx, mux, r.service)
	if err != nil {
		return err
	}
	//>>
	err = mux.HandlePath(
		http.MethodGet,
		"/doc.html",
		func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			b, err := proto.DocHtml.ReadFile("doc.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			_, err = w.Write(b)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)

			return
		},
	)
	if err != nil {
		return errors.Wrap(err, "handling doc")
	}
	err = mux.HandlePath(
		http.MethodGet,
		"/auth.swagger.json",
		func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			b, err := proto.SwaggerJson.ReadFile("auth.swagger.json")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			_, err = w.Write(b)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)

			return
		},
	)
	if err != nil {
		return errors.Wrap(err, "handling doc")
	}
	//<<
	err = mux.HandlePath(
		http.MethodOptions,
		"/*",
		func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

			w.WriteHeader(http.StatusNoContent)

			return
		},
	)
	if err != nil {
		return errors.Wrap(err, "handling http cors")
	}

	r.server = &http.Server{
		//ReadTimeout:  5 * time.Second,
		//WriteTimeout: 5 * time.Second,
		//IdleTimeout:  5 * time.Second,
		Addr:    r.httpAddr,
		Handler: NewHandler(mux),
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	return r.server.ListenAndServe()
}

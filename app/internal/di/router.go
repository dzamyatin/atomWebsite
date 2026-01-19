package di

import (
	"context"
	"net/http"

	atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"
	"github.com/dzamyatin/atomWebsite/internal/grpc/grpc"
	"github.com/dzamyatin/atomWebsite/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func newHttpRouter(
	serviceAuth grpc.AuthServer,
	serviceShop grpc.ShopServer,
) *grpc.HttpRouter {
	return grpc.NewHttpRouter(
		func(ctx context.Context, mux *runtime.ServeMux) error {
			return atomWebsite.RegisterShopHandlerServer(ctx, mux, serviceShop)
		},
		func(ctx context.Context, mux *runtime.ServeMux) error {
			return atomWebsite.RegisterAuthHandlerServer(ctx, mux, serviceAuth)
		},
		func(ctx context.Context, mux *runtime.ServeMux) error {
			return mux.HandlePath(
				http.MethodGet,
				"/doc.html",
				func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
					w.WriteHeader(http.StatusOK)

					b, err := proto.DocHtml.ReadFile("doc.html")
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					_, err = w.Write(b)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					return
				},
			)
		},
		func(ctx context.Context, mux *runtime.ServeMux) error {
			return mux.HandlePath(
				http.MethodGet,
				"/auth.swagger.json",
				func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
					w.WriteHeader(http.StatusOK)

					b, err := proto.SwaggerJson.ReadFile("auth.swagger.json")
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					_, err = w.Write(b)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					return
				},
			)
		},
		func(ctx context.Context, mux *runtime.ServeMux) error {
			return mux.HandlePath(
				http.MethodOptions,
				"/*",
				func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

					w.WriteHeader(http.StatusNoContent)

					return
				},
			)
		},
	)
}

package httpserver

import (
	"net/http"
)

type CorsHandlerMiddleware struct {
	h    http.Handler
	host string
}

func NewCorsHandlerMiddleware(
	h http.Handler,
	host string,
) *CorsHandlerMiddleware {
	return &CorsHandlerMiddleware{
		h:    h,
		host: host,
	}
}

func (h *CorsHandlerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("access-control-allow-credentials", "true")
	w.Header().Set("access-control-allow-headers", "Content-Type, Authorization")
	w.Header().Set("access-control-allow-methods", "PUT, GET, POST, PATCH, DELETE, OPTIONS")
	w.Header().Set("access-control-allow-origin", h.host)
	w.Header().Set("access-control-max-age", "1728000")

	h.h.ServeHTTP(w, r)
}

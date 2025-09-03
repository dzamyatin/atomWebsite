package grpc

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
)

type MuxCustomization func(ctx context.Context, mux *runtime.ServeMux) error

type HttpRouter struct {
	muxCustomizations []MuxCustomization
}

func NewHttpRouter(muxCustomizations ...MuxCustomization) *HttpRouter {
	return &HttpRouter{muxCustomizations: muxCustomizations}
}

func (r *HttpRouter) Add(customization MuxCustomization) {
	r.muxCustomizations = append(r.muxCustomizations, customization)
}

func (r *HttpRouter) Apply(ctx context.Context, mux *runtime.ServeMux) error {
	for _, customization := range r.muxCustomizations {
		if err := customization(ctx, mux); err != nil {
			return errors.Wrap(err, "failed to apply customization")
		}
	}

	return nil
}

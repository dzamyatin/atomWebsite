package grpc

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
)

type MuxCustomization func(mux *runtime.ServeMux) error

type HttpRouter struct {
	muxCustomizations []MuxCustomization
}

func NewHttpRouter(muxCustomizations ...MuxCustomization) *HttpRouter {
	return &HttpRouter{muxCustomizations: muxCustomizations}
}

func (r *HttpRouter) Add(customization MuxCustomization) {
	r.muxCustomizations = append(r.muxCustomizations, customization)
}

func (r *HttpRouter) Apply(mux *runtime.ServeMux) error {
	for _, customization := range r.muxCustomizations {
		if err := customization(mux); err != nil {
			return errors.Wrap(err, "failed to apply customization")
		}
	}

	return nil
}

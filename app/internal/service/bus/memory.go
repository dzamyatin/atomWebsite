package bus

import (
	"context"
	"github.com/pkg/errors"
)

const BusMemory BusName = "memory"

type MemoryBus struct {
	BaseBus
}

func NewMemoryBus() *MemoryBus {
	return &MemoryBus{
		BaseBus{
			dispatcher: make(map[string]IHandler),
		},
	}
}

func (r *MemoryBus) Dispatch(ctx context.Context, command ICommand) error {
	handler, err := r.GetHandler(command)

	if err != nil {
		return errors.Wrap(err, "get handler")
	}

	return errors.Wrap(handler.Handle(ctx, command), "execute")
}

package bus

import (
	"context"
	"github.com/pkg/errors"
)

const BusMemory BusName = "memory"

type MemoryBus struct {
	dispatcher map[string]IHandler
}

func NewMemoryBus() *MemoryBus {
	return &MemoryBus{dispatcher: make(map[string]IHandler)}
}

func (r *MemoryBus) Register(command ICommand, handler IHandler, name BusName) error {
	r.dispatcher[command.GetName()] = handler

	return nil
}

func (r *MemoryBus) Dispatch(ctx context.Context, command ICommand) error {
	handler, ok := r.dispatcher[command.GetName()]
	if !ok {
		return errors.New("no Handler found")
	}

	return errors.Wrap(handler.Handle(ctx, command), "execute")
}

func (r *MemoryBus) GetName() BusName {
	return BusMemory
}

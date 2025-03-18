package di

import (
	"github.com/dzamyatin/atomWebsite/internal/service/bus"
	"github.com/dzamyatin/atomWebsite/internal/service/command"
)

func newHandlerRegistry(
	handler *command.RegisterHandler,
) bus.HandlerRegistry {
	return bus.HandlerRegistry{
		{
			Command: new(command.RegisterCommand),
			Handler: handler,
			BusName: bus.BusMemory,
		},
	}
}

func newBus(
	memoryBus *bus.MemoryBus,
	registry bus.HandlerRegistry,
) *bus.MainBus {
	return bus.NewBus(
		map[bus.BusName]bus.IBus{
			bus.BusMemory: memoryBus,
		},
		registry,
	)
}

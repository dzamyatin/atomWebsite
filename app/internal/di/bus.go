package di

import (
	"github.com/dzamyatin/atomWebsite/internal/service/bus"
	"github.com/dzamyatin/atomWebsite/internal/service/command"
	"github.com/dzamyatin/atomWebsite/internal/service/db"
	"go.uber.org/zap"
	"time"
)

func newHandlerRegistry(
	handler *command.RegisterHandler,
) bus.HandlerRegistry {
	return bus.HandlerRegistry{
		{
			Command: new(command.RegisterCommand),
			Handler: handler,
			//BusName: bus.BusMemory,
			BusName: bus.BusPostgres,
		},
	}
}

func newPostgresBus(
	db db.IDatabase,
	logger *zap.Logger,
) *bus.PostgresBus {
	return bus.NewPostgresBus(
		"main",
		db,
		30*time.Second,
		5,
		logger,
	)
}

func newBus(
	postgresBus *bus.PostgresBus,
	memoryBus *bus.MemoryBus,
	registry bus.HandlerRegistry,
	logger *zap.Logger,
) *bus.MainBus {
	return bus.NewBus(
		map[bus.BusName]bus.IBus{
			bus.BusMemory:   memoryBus,
			bus.BusPostgres: postgresBus,
		},
		registry,
		logger,
	)
}

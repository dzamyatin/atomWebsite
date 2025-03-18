package bus

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type MainBus struct {
	buses    map[BusName]IBus
	commands map[string][]BusName
	logger   *zap.Logger
}

func NewBus(
	buses map[BusName]IBus,
	handlerRegistry HandlerRegistry,
	logger *zap.Logger,
) *MainBus {
	mappedBuses := make(map[BusName]IBus)

	for name, bus := range buses {
		mappedBuses[name] = bus
	}

	bus := &MainBus{
		buses:    mappedBuses,
		commands: map[string][]BusName{},
		logger:   logger,
	}

	for _, handler := range handlerRegistry {
		err := bus.Register(handler.Command, handler.Handler, handler.BusName)
		if err != nil {
			panic(err)
		}
	}

	return bus
}

func (b *MainBus) Register(command ICommand, handler IHandler, name BusName) error {
	if command.GetName() == "" {
		return errors.New("invalid Command name")
	}

	b.commands[command.GetName()] = append(b.commands[command.GetName()], name)

	bus, ok := b.buses[name]
	if !ok {
		return errors.Errorf("bus not registered %v", name)
	}

	return bus.Register(command, handler, name)
}

func (b *MainBus) Dispatch(ctx context.Context, command ICommand) error {
	buses, ok := b.commands[command.GetName()]

	if !ok {
		return errors.New("bus not found across commands")
	}

	var err error
	for _, bus := range buses {
		var busInstance IBus
		busInstance, ok = b.buses[bus]

		if !ok {
			return errors.New("bus not found across available busses")
		}

		err = busInstance.Dispatch(ctx, command)

		if err != nil {
			b.logger.Error(
				"Error while dispatchung message through bus",
				zap.Any("Command", command),
				zap.String("bus", string(bus)),
				zap.Error(err),
			)
		}
	}

	return err
}

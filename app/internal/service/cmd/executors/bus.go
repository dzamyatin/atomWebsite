package executors

import (
	"context"

	mainarg "github.com/dzamyatin/atomWebsite/internal/service/arg"
	"github.com/dzamyatin/atomWebsite/internal/service/bus"
	"github.com/dzamyatin/atomWebsite/internal/service/process"
	"go.uber.org/zap"
)

type ArgBusProcess struct {
	mainarg.Arg
	QueueName string `arg:"-q,required" help:"Name of queue to process"`
}

type BusProcessCommand struct {
	logger         *zap.Logger
	bus            *bus.PostgresBus
	processManager *process.ProcessShutdownerManager
}

func NewBusProcessCommand(
	logger *zap.Logger,
	bus *bus.PostgresBus,
	_ *bus.MainBus, //to initialize handlers throgh di
	processManager *process.ProcessShutdownerManager,
) *BusProcessCommand {
	return &BusProcessCommand{
		logger:         logger,
		bus:            bus,
		processManager: processManager,
	}
}

func (r *BusProcessCommand) Execute(ctx context.Context, u ArgBusProcess) error {
	return r.processManager.Run(
		ctx,
		process.Process{
			Name: "bus-process",
			Object: process.NewProcessor(
				func(ctx context.Context) error {
					return r.bus.ProcessCycle(ctx, u.QueueName)
				},
				func() error {
					return nil
				},
			),
		},
	)
}

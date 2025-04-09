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
	logger *zap.Logger
	bus    *bus.PostgresBus
}

func NewBusProcessCommand(
	logger *zap.Logger,
	bus *bus.PostgresBus,
) *BusProcessCommand {
	return &BusProcessCommand{logger: logger, bus: bus}
}

func (r *BusProcessCommand) Execute(ctx context.Context, u ArgBusProcess) error {
	return process.NewProcessManager(
		r.logger,
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
	).Start(ctx)
}

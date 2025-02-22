package cmd

import (
	"context"

	"github.com/dzamyatin/atomWebsite/internal/di"
	"github.com/dzamyatin/atomWebsite/internal/service/cmd/executors"
)

const defaultHelp = "There is no special conditions"

func init() {
	GetRegistry().Register(
		"grpc",
		NewCommand[executors.ArgGrpcProcess](
			func(ctx context.Context, arg executors.ArgGrpcProcess) IExecuter[executors.ArgGrpcProcess] {
				v, err := di.InitializeGrpcProcessCommand(
					ctx,
					arg.GetParsedConfig(),
				)
				if err != nil {
					panic(err)
				}

				return v
			},
			defaultHelp,
		),
	)
	GetRegistry().Register(
		"bot",
		NewCommand[executors.ArgTelegramBotProcess](
			func(ctx context.Context, _ executors.ArgTelegramBotProcess) IExecuter[executors.ArgTelegramBotProcess] {
				v, err := di.InitializeTelegramBotProcessCommand(ctx)
				if err != nil {
					panic(err)
				}

				return v
			},
			defaultHelp,
		),
	)
	GetRegistry().Register(
		"bus",
		NewCommand[executors.ArgBusProcess](
			func(ctx context.Context, _ executors.ArgBusProcess) IExecuter[executors.ArgBusProcess] {
				v, err := di.InitializeBusProcessCommand(ctx)
				if err != nil {
					panic(err)
				}

				return v
			},
			defaultHelp,
		),
	)
	GetRegistry().Register(
		"migration-create",
		NewCommand[executors.ArgMigrationCreate](
			func(ctx context.Context, _ executors.ArgMigrationCreate) IExecuter[executors.ArgMigrationCreate] {
				return di.InitializeMigrationCreateCommand(ctx)
			},
			defaultHelp,
		),
	)
	GetRegistry().Register(
		"migration-down",
		NewCommand[executors.ArgMigrationDown](
			func(ctx context.Context, _ executors.ArgMigrationDown) IExecuter[executors.ArgMigrationDown] {
				v, err := di.InitializeMigrationDownCommand(ctx)

				if err != nil {
					panic(err)
				}

				return v
			},
			defaultHelp,
		),
	)

	GetRegistry().Register(
		"migration-up",
		NewCommand[executors.ArgMogrationUp](
			func(ctx context.Context, _ executors.ArgMogrationUp) IExecuter[executors.ArgMogrationUp] {
				v, err := di.InitializeMigrationUpCommand(ctx)

				if err != nil {
					panic(err)
				}

				return v
			},
			defaultHelp,
		),
	)
}

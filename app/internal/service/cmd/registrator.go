package cmd

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/di"
	"github.com/dzamyatin/atomWebsite/internal/service/cmd/executors"
)

const defaultHelp = "There is no special conditions"

func init() {
	GetRegistry().Register(
		"migration-create",
		NewCommand[executors.ArgMigrationCreate](
			func(ctx context.Context) IExecuter[executors.ArgMigrationCreate] {
				return di.InitializeMigrationCreateCommand(ctx)
			},
			defaultHelp,
		),
	)
	GetRegistry().Register(
		"migration-down",
		NewCommand[executors.ArgMigrationDown](
			func(ctx context.Context) IExecuter[executors.ArgMigrationDown] {
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
			func(ctx context.Context) IExecuter[executors.ArgMogrationUp] {
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

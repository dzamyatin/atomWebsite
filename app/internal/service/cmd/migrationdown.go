package cmd

import (
	"context"
	"github.com/alexflint/go-arg"
	"github.com/dzamyatin/atomWebsite/internal/di"
	"go.uber.org/zap"
)

func init() {
	GetRegistry().Register("migration-down", NewMigrationDownCommand())
}

type MigrationDownCommand struct{}

func (r *MigrationDownCommand) Execute(ctx context.Context) int {
	args := &ArgMogration{}
	arg.MustParse(args)
	err := di.CreateConfig(args.Config)
	logger := di.InitializeLogger(ctx)

	if err != nil {
		logger.Error("could not create config", zap.Error(err))

		panic(err)
	}

	com, err := di.InitializeMigrationDownCommand(ctx)

	if err != nil {
		logger.Error("could not initialize migration down command", zap.Error(err))

		panic(err)
	}

	err = com.Execute(ctx)

	if err != nil {
		logger.Error("could not execute migration down command", zap.Error(err))

		panic(err)
	}

	return 0
}

func NewMigrationDownCommand() *MigrationDownCommand {
	return &MigrationDownCommand{}
}

func (r *MigrationDownCommand) Help() string {
	return "There is no special requirements"
}

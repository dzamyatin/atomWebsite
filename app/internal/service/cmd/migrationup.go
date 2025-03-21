package cmd

import (
	"context"
	"github.com/alexflint/go-arg"
	"github.com/dzamyatin/atomWebsite/internal/di"
	mainarg "github.com/dzamyatin/atomWebsite/internal/service/arg"
	"go.uber.org/zap"
)

func init() {
	GetRegistry().Register("migration-up", NewMigrationUpCommand())
}

type ArgMogration struct {
	mainarg.Arg
}

type MigrationUpCommand struct{}

func (r *MigrationUpCommand) Execute(ctx context.Context) int {
	args := &ArgMogration{}
	arg.MustParse(args)
	err := di.CreateConfig(args.Config)
	logger := di.InitializeLogger(ctx)

	if err != nil {
		logger.Error("could not create config", zap.Error(err))

		panic(err)
	}

	com, err := di.InitializeMigrationUpCommand(ctx)

	if err != nil {
		logger.Error("could not initialize migration up command", zap.Error(err))

		panic(err)
	}

	err = com.Execute(ctx)

	if err != nil {
		logger.Error("could not execute migration up command", zap.Error(err))

		panic(err)
	}

	return 0
}

func NewMigrationUpCommand() *MigrationUpCommand {
	return &MigrationUpCommand{}
}

func (r *MigrationUpCommand) Help() string {
	return "There is no special requirements"
}

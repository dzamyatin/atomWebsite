package command

import (
	"context"
	"github.com/alexflint/go-arg"
	"github.com/dzamyatin/atomWebsite/internal/di"
	mainarg "github.com/dzamyatin/atomWebsite/internal/service/arg"
	"go.uber.org/zap"
)

func init() {
	GetRegistry().Register("migration-up", NewMigrationCommand())
}

type ArgMogration struct {
	mainarg.Arg
}

type MigrationCommand struct{}

func (r *MigrationCommand) Execute(ctx context.Context) int {
	args := &ArgMogration{}
	arg.MustParse(args)
	err := di.CreateConfig(args.Config)
	logger := di.InitializeLogger()

	if err != nil {
		logger.Error("could not create config", zap.Error(err))

		panic(err)
	}

	com, err := di.InitializeMigrationUpCommand()

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

func NewMigrationCommand() *MigrationCommand {
	return &MigrationCommand{}
}

func (r *MigrationCommand) Help() string {
	return "There is no special requirements"
}

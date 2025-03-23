package executors

import (
	"context"
	mainarg "github.com/dzamyatin/atomWebsite/internal/service/arg"
	usecasemigration "github.com/dzamyatin/atomWebsite/internal/usecase/migration"
	"go.uber.org/zap"
)

type MigrationDownCommand struct {
	logger *zap.Logger
	com    *usecasemigration.Down
}

type ArgMigrationDown struct {
	mainarg.Arg
}

func NewMigrationDownCommand(logger *zap.Logger, com *usecasemigration.Down) *MigrationDownCommand {
	return &MigrationDownCommand{logger: logger, com: com}
}

func (r *MigrationDownCommand) Execute(ctx context.Context, _ ArgMigrationDown) error {
	err := r.com.Execute(ctx)

	if err != nil {
		r.logger.Error("could not execute migration down command", zap.Error(err))

		panic(err)
	}

	return nil
}

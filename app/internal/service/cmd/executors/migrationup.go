package executors

import (
	"context"
	mainarg "github.com/dzamyatin/atomWebsite/internal/service/arg"
	usecasemigration "github.com/dzamyatin/atomWebsite/internal/usecase/migration"
	"go.uber.org/zap"
)

type ArgMogrationUp struct {
	mainarg.Arg
}

type MigrationUpCommand struct {
	logger *zap.Logger
	com    *usecasemigration.Up
}

func NewMigrationUpCommand(logger *zap.Logger, com *usecasemigration.Up) *MigrationUpCommand {
	return &MigrationUpCommand{logger: logger, com: com}
}

func (r *MigrationUpCommand) Execute(ctx context.Context, _ ArgMogrationUp) error {
	err := r.com.Execute(ctx)

	if err != nil {
		r.logger.Error("could not execute migration up command", zap.Error(err))

		panic(err)
	}

	return nil
}

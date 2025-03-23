package executors

import (
	"context"
	mainarg "github.com/dzamyatin/atomWebsite/internal/service/arg"
	"go.uber.org/zap"
)

type ArgMigrationCreate struct {
	mainarg.Arg
	Name string `arg:"-n,required" help:"Name of migration"`
}

type MigrationCreateCommand struct {
	logger *zap.Logger
}

func NewMigrationCreateCommand(logger *zap.Logger) *MigrationCreateCommand {
	return &MigrationCreateCommand{logger: logger}
}

func (m *MigrationCreateCommand) Execute(ctx context.Context, u ArgMigrationCreate) error {
	return nil
}

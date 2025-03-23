//go:build wireinject
// +build wireinject

package di

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/service/cmd/executors"

	"github.com/google/wire"
)

func InitializeMigrationUpCommand(ctx context.Context) (*executors.MigrationUpCommand, error) {
	wire.Build(set)

	return &executors.MigrationUpCommand{}, nil
}

func InitializeMigrationCreateCommand(ctx context.Context) *executors.MigrationCreateCommand {
	wire.Build(set)

	return &executors.MigrationCreateCommand{}
}

func InitializeMigrationDownCommand(ctx context.Context) (*executors.MigrationDownCommand, error) {
	wire.Build(set)

	return &executors.MigrationDownCommand{}, nil
}

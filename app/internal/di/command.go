//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"github.com/dzamyatin/atomWebsite/internal/service/cmd/executors"
	"github.com/dzamyatin/atomWebsite/internal/service/config"

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

func InitializeBusProcessCommand(ctx context.Context) (*executors.BusProcessCommand, error) {
	wire.Build(set)

	return &executors.BusProcessCommand{}, nil
}

func InitializeTelegramBotProcessCommand(ctx context.Context) (*executors.TelegramBotProcessCommand, error) {
	wire.Build(set)

	return &executors.TelegramBotProcessCommand{}, nil
}

func InitializeGrpcProcessCommand(ctx context.Context, cfg config.AppConfig) (*executors.GrpcProcessCommand, error) {
	wire.Build(set)

	return &executors.GrpcProcessCommand{}, nil
}

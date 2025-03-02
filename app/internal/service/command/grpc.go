package command

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/di"
	"github.com/dzamyatin/atomWebsite/internal/service/arg"
	"log"
)

func init() {
	GetRegistry().Register("grpc", NewGRPCCommand())
}

type GRPCCommand struct{}

func (G GRPCCommand) Execute(ctx context.Context) int {
	a := arg.MustNewArg()
	err := di.CreateConfig(a.Config)

	if err != nil {
		log.Fatalf("failed to create config: %v", err)
	}

	manager, err := di.InitializeGRPCProcessManager()
	if err != nil {
		log.Fatalf("failed to initialize grpc process manager: %v", err)
	}

	err = manager.Start(ctx)

	return 0
}

func (G GRPCCommand) Help() string {
	return "Usage: grpc --config config.yaml"
}

func NewGRPCCommand() *GRPCCommand {
	return &GRPCCommand{}
}

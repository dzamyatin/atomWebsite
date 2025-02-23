//go:build wireinject
// +build wireinject

package di

import (
	grpcservice "github.com/dzamyatin/atomWebsite/internal/service/grpc"
	"github.com/dzamyatin/atomWebsite/internal/service/process"
	"github.com/google/wire"
	_ "github.com/lib/pq"
)

//go:generate go tool wire

func InitializeGRPCProcessManager() *process.ProcessManager {
	wire.Build(
		newGRPCProcessManager,
		newLogger,
		newServer,
		newGrpcServer,
		grpcservice.NewAuthServer,
		process.NewSignalListener,
	)

	return &process.ProcessManager{}
}

//go:build wireinject
// +build wireinject

package di

import (
	"github.com/dzamyatin/atomWebsite/internal/grpc/grpc"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	"github.com/dzamyatin/atomWebsite/internal/service/process"
	"github.com/dzamyatin/atomWebsite/internal/usecase"
	"github.com/google/wire"
	_ "github.com/lib/pq"
)

//go:generate go tool wire

func InitializeGRPCProcessManager() (*process.ProcessManager, error) {
	wire.Build(
		newGRPCProcessManager,
		newLogger,
		newServer,
		newGrpcServer,
		grpc.NewAuthServer,
		process.NewSignalListener,
		usecase.NewRegistrationUseCase,
		repository.NewUserRepository,
		wire.Bind(new(repository.IUserRepository), new(*repository.UserRepository)),
		newDb,
	)

	return &process.ProcessManager{}, nil
}

//go:build wireinject
// +build wireinject

package di

import (
	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/dzamyatin/atomWebsite/internal/grpc/grpc"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	"github.com/dzamyatin/atomWebsite/internal/service/process"
	userservice "github.com/dzamyatin/atomWebsite/internal/service/user"
	"github.com/dzamyatin/atomWebsite/internal/usecase"
	usecasemigration "github.com/dzamyatin/atomWebsite/internal/usecase/migration"
	"github.com/dzamyatin/atomWebsite/internal/validator"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

//go:generate go tool wire

var set = wire.NewSet(
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
	wire.Bind(new(entity.PasswordEncoder), new(*userservice.PasswordEncoder)),
	wire.Bind(new(entity.PasswordComparator), new(*userservice.PasswordEncoder)),
	userservice.NewPasswordEncoder,
	wire.Bind(new(validator.IRegistrationValidator), new(*validator.RegistrationValidator)),
	validator.NewRegistrationValidator,
	usecasemigration.NewUp,
	usecasemigration.NewDown,
)

func InitializeGRPCProcessManager() (*process.ProcessManager, error) {
	wire.Build(set)

	return &process.ProcessManager{}, nil
}

func InitializeMigrationUpCommand() (*usecasemigration.Up, error) {
	wire.Build(set)

	return &usecasemigration.Up{}, nil
}

func InitializeMigrationDownCommand() (*usecasemigration.Down, error) {
	wire.Build(set)

	return &usecasemigration.Down{}, nil
}

func InitializeLogger() *zap.Logger {
	wire.Build(set)

	return &zap.Logger{}
}

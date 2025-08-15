//go:build wireinject
// +build wireinject

package di

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/dzamyatin/atomWebsite/internal/grpc/grpc"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	serviceauth "github.com/dzamyatin/atomWebsite/internal/service/auth"
	"github.com/dzamyatin/atomWebsite/internal/service/bus"
	"github.com/dzamyatin/atomWebsite/internal/service/cmd/executors"
	"github.com/dzamyatin/atomWebsite/internal/service/db"
	"github.com/dzamyatin/atomWebsite/internal/service/handler"
	servicemessengersender "github.com/dzamyatin/atomWebsite/internal/service/messenger/sender"
	servicemessengerstatemachine "github.com/dzamyatin/atomWebsite/internal/service/messenger/statemachine"
	servicemessengerstatemachinestate "github.com/dzamyatin/atomWebsite/internal/service/messenger/statemachine/state"
	"github.com/dzamyatin/atomWebsite/internal/service/metric"
	"github.com/dzamyatin/atomWebsite/internal/service/process"
	"github.com/dzamyatin/atomWebsite/internal/service/time"
	userservice "github.com/dzamyatin/atomWebsite/internal/service/user"
	"github.com/dzamyatin/atomWebsite/internal/usecase"
	usecasemessenger "github.com/dzamyatin/atomWebsite/internal/usecase/messenger"
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
	usecase.NewRegistration,
	repository.NewUserRepository,
	wire.Bind(new(repository.IUserRepository), new(*repository.UserRepository)),
	newDb,
	newDbx,
	db.NewDatabase,
	wire.Bind(new(db.IDatabase), new(*db.Database)),
	wire.Bind(new(entity.PasswordEncoder), new(*userservice.PasswordEncoder)),
	wire.Bind(new(entity.PasswordComparator), new(*userservice.PasswordEncoder)),
	userservice.NewPasswordEncoder,
	usecasemigration.NewUp,
	usecasemigration.NewDown,
	metric.NewMetric,
	metric.NewRegistry,
	usecase.NewLogin,
	wire.Bind(new(serviceauth.IProvider), new(*serviceauth.SequentialProvider)),
	newSequentialProvider,
	wire.Bind(new(serviceauth.IJWT), new(*serviceauth.JWT)),
	newJWT,
	wire.Bind(new(bus.IBus), new(*bus.MainBus)),
	newBus,
	newHandlerRegistry,
	bus.NewMemoryBus,
	handler.NewRegisterHandler,
	executors.NewMigrationCreateCommand,
	executors.NewMigrationDownCommand,
	executors.NewMigrationUpCommand,
	newPostgresBus,
	executors.NewBusProcessCommand,
	newHTTPServer,
	newMailer,
	servicetime.NewTime,
	wire.Bind(new(servicetime.ITime), new(*servicetime.Time)),
	repository.NewRandomizerRepository,
	wire.Bind(new(repository.IRandomizerRepository), new(*repository.RandomizerRepository)),
	usecase.NewConfirmEmailUseCase,
	usecase.NewSendEmailConfirmationUseCase,
	//newMessenger,
	usecase.NewRememberPasswordUseCase,
	validator.NewValidator,
	usecase.NewChangePasswordUseCase,
	newTelegramBotServer,
	executors.NewTelegramBotProcessCommand,
	wire.Bind(new(repository.IChatRepository), new(*repository.ChatRepository)),
	repository.NewChatRepository,
	servicemessengerstatemachinestate.NewInitialState,
	servicemessengerstatemachinestate.NewWaitForPhone,
	newStateRegistry,
	usecasemessenger.NewReceiveMessageUseCase,
	servicemessengerstatemachine.NewStateMachineFactory,
	newMessengerServerRegistry,
	repository.NewSenderRepository,
	wire.Bind(new(repository.ISenderRepository), new(*repository.SenderRepository)),
	servicemessengerstatemachinestate.NewPhoneStoredState,
	repository.NewCounterRepository,
	wire.Bind(new(repository.ICounterRepository), new(*repository.CounterRepository)),
	servicemessengersender.NewSequentiallySender,
	wire.Bind(new(servicemessengersender.ISenderService), new(*servicemessengersender.SequentiallySender)),
)

func InitializeGRPCProcessManager(ctx context.Context) (*process.ProcessManager, error) {
	wire.Build(set)

	return &process.ProcessManager{}, nil
}

func InitializeLogger(ctx context.Context) *zap.Logger {
	wire.Build(set)

	return &zap.Logger{}
}

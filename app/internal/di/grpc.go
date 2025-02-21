//go:build wireinject
// +build wireinject

package di

import (
	"fmt"
	atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"
	grpcservice "github.com/dzamyatin/atomWebsite/internal/service/grpc"
	"github.com/dzamyatin/atomWebsite/internal/service/process"
	"github.com/google/wire"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
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

func newGRPCProcessManager(
	logger *zap.Logger,
	serv *grpcservice.GRPCServer,
	listener *process.SignalListener,
) *process.ProcessManager {
	return process.NewProcessManager(
		logger,
		process.Process{
			Name:   "grpc server",
			Object: serv,
		},
		process.Process{
			Name:   "signal listener",
			Object: listener,
		},
	)
}

func newLogger() *zap.Logger {
	return zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.AddSync(os.Stdout),
			zapcore.InfoLevel,
		),
	)
}

func newServer(
	grpcServer *grpc.Server,
) *grpcservice.GRPCServer {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8502))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return grpcservice.NewGRPCServer(
		lis,
		grpcServer,
	)
}

func newGrpcServer(server grpcservice.AuthServer) *grpc.Server {
	grpcServer := grpc.NewServer()

	atomWebsite.RegisterAuthServer(grpcServer, server)

	return grpcServer
}

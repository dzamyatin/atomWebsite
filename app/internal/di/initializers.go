package di

import (
	"database/sql"
	"fmt"
	atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"
	grpcservice2 "github.com/dzamyatin/atomWebsite/internal/grpc/grpc"
	"github.com/dzamyatin/atomWebsite/internal/service/process"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/url"
	"os"
)

func newGRPCProcessManager(
	logger *zap.Logger,
	serv *grpcservice2.GRPCServer,
	listener *process.SignalListener,
	db *sql.DB,
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
		process.Process{
			Name: "database shutdowner",
			Object: process.NewShutdowner(
				logger,
				func() error {
					return db.Close()
				},
			),
		},
	)
}

func newDb() (*sql.DB, error) {
	config := getConfig().Db

	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(config.Username, config.Password),
		Host:   config.Host + ":" + config.Port,
	}

	db, err := sql.Open("postgres", u.String())

	return db, err
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
) *grpcservice2.GRPCServer {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8502))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	return grpcservice2.NewGRPCServer(
		lis,
		grpcServer,
	)
}

func newGrpcServer(server grpcservice2.AuthServer) *grpc.Server {
	grpcServer := grpc.NewServer()

	atomWebsite.RegisterAuthServer(grpcServer, server)

	return grpcServer
}

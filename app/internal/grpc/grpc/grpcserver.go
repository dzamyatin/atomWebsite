package grpc

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type GRPCServer struct {
	logger      *zap.Logger
	netListener net.Listener
	grpcServer  *grpc.Server
}

func NewGRPCServer(
	logger *zap.Logger,
	netListener net.Listener,
	grpcServer *grpc.Server,
) *GRPCServer {

	return &GRPCServer{
		logger:      logger,
		netListener: netListener,
		grpcServer:  grpcServer,
	}
}

func (s *GRPCServer) Shutdown() error {
	s.grpcServer.GracefulStop()

	return nil
}

func (s *GRPCServer) Start(_ context.Context) error {
	s.logger.Info("starting gRPC server", zap.String("address", s.netListener.Addr().String()))
	return s.grpcServer.Serve(s.netListener)
}

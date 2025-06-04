package grpc

import (
	"context"
	"google.golang.org/grpc"
	"net"
)

type GRPCServer struct {
	netListener net.Listener
	grpcServer  *grpc.Server
}

func NewGRPCServer(
	netListener net.Listener,
	grpcServer *grpc.Server,
) *GRPCServer {

	return &GRPCServer{
		netListener: netListener,
		grpcServer:  grpcServer,
	}
}

func (s *GRPCServer) Shutdown() error {
	s.grpcServer.GracefulStop()

	return nil
}

func (s *GRPCServer) Start(_ context.Context) error {
	return s.grpcServer.Serve(s.netListener)
}

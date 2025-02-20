package main

import (
	"fmt"
	atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"
	"google.golang.org/grpc"
	"log"
	"net"
)

type AuthServer struct {
	atomWebsite.UnimplementedAuthServer
}

func main() {
	fmt.Println("Hello world!")

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8502))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	atomWebsite.RegisterAuthServer(grpcServer, AuthServer{})

	grpcServer.Serve(lis)

	fmt.Println("Done")
}

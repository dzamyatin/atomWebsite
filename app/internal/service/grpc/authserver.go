package grpcservice

import atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"

type AuthServer struct {
	atomWebsite.UnimplementedAuthServer
}

func NewAuthServer() AuthServer {
	return AuthServer{}
}

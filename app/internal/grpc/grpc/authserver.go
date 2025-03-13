package grpc

import (
	"context"
	atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"
	"github.com/dzamyatin/atomWebsite/internal/request"
	"github.com/dzamyatin/atomWebsite/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	atomWebsite.UnimplementedAuthServer
	registerUseCase *usecase.Registration
}

func NewAuthServer(registerUseCase *usecase.Registration) AuthServer {
	return AuthServer{
		registerUseCase: registerUseCase,
	}
}

func (r AuthServer) Register(ctx context.Context, req *atomWebsite.RegisterRequest) (*atomWebsite.RegisterResponse, error) {
	err := r.registerUseCase.Execute(
		ctx,
		request.RegistrationRequest{
			Email:    req.Email,
			Password: req.Password,
			Phone:    req.Phone,
		},
	)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Registration fail: %s", err.Error())
	}

	return &atomWebsite.RegisterResponse{}, nil
}

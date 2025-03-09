package grpc

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/dto"
	atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"
	"github.com/dzamyatin/atomWebsite/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	atomWebsite.UnimplementedAuthServer
	registerUseCase *usecase.RegistrationUseCase
}

func NewAuthServer(registerUseCase *usecase.RegistrationUseCase) AuthServer {
	return AuthServer{
		registerUseCase: registerUseCase,
	}
}

func (r AuthServer) Register(ctx context.Context, req *atomWebsite.RegisterRequest) (*atomWebsite.RegisterResponse, error) {
	err := r.registerUseCase.Execute(
		ctx,
		dto.RegistrationRequest{
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

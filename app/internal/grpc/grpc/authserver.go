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
	loginUseCase    *usecase.Login
}

func NewAuthServer(registerUseCase *usecase.Registration, loginUseCase *usecase.Login) AuthServer {
	return AuthServer{
		registerUseCase: registerUseCase,
		loginUseCase:    loginUseCase,
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

func (r AuthServer) Login(ctx context.Context, req *atomWebsite.LoginRequest) (*atomWebsite.LoginResponse, error) {
	res, err := r.loginUseCase.Execute(ctx, request.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
	})

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Login fail: %s", err.Error())
	}

	return &atomWebsite.LoginResponse{
		Token: res.Token,
	}, nil
}

package grpc

import (
	"context"
	atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"
	"github.com/dzamyatin/atomWebsite/internal/request"
	"github.com/dzamyatin/atomWebsite/internal/service/bus"
	"github.com/dzamyatin/atomWebsite/internal/usecase"
	"github.com/guregu/null/v6"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	atomWebsite.UnimplementedAuthServer
	bus             bus.IBus
	registerUseCase *usecase.Registration
	loginUseCase    *usecase.Login
}

func NewAuthServer(
	registerUseCase *usecase.Registration,
	loginUseCase *usecase.Login,
	bus bus.IBus,
) AuthServer {
	return AuthServer{
		registerUseCase: registerUseCase,
		loginUseCase:    loginUseCase,
		bus:             bus,
	}
}

func (r AuthServer) ConfirmEmail(
	context.Context,
	*atomWebsite.ConfirmEmailRequest,
) (*atomWebsite.ConfirmEmailResponse, error) {
	return nil, nil
}

func (r AuthServer) Register(ctx context.Context, req *atomWebsite.RegisterRequest) (*atomWebsite.RegisterResponse, error) {
	err := r.registerUseCase.Execute(
		ctx,
		request.RegistrationRequest{
			Email:    null.NewValue(req.Email, req.Email != ""),
			Password: req.Password,
			Phone:    null.NewValue(req.Phone, req.Phone != ""),
		},
	)

	//err := r.bus.Dispatch(
	//	ctx,
	//	&command.RegisterCommand{Req: request.RegistrationRequest{
	//		Email:    req.Email,
	//		Password: req.Password,
	//		Phone:    req.Phone,
	//	}},
	//)

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

package grpc

import (
	"context"
	atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"
	"github.com/dzamyatin/atomWebsite/internal/request"
	"github.com/dzamyatin/atomWebsite/internal/service/bus"
	"github.com/dzamyatin/atomWebsite/internal/usecase"
	"github.com/guregu/null/v6"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	atomWebsite.UnimplementedAuthServer
	bus                          bus.IBus
	registerUseCase              *usecase.Registration
	loginUseCase                 *usecase.Login
	confirmEmailUseCase          *usecase.ConfirmEmailUseCase
	sendEmailConfirmationUseCase *usecase.SendEmailConfirmationUseCase
}

func NewAuthServer(
	registerUseCase *usecase.Registration,
	loginUseCase *usecase.Login,
	confirmEmailUseCase *usecase.ConfirmEmailUseCase,
	bus bus.IBus,
	sendEmailConfirmationUseCase *usecase.SendEmailConfirmationUseCase,
) AuthServer {
	return AuthServer{
		registerUseCase:              registerUseCase,
		loginUseCase:                 loginUseCase,
		bus:                          bus,
		confirmEmailUseCase:          confirmEmailUseCase,
		sendEmailConfirmationUseCase: sendEmailConfirmationUseCase,
	}
}

func (r AuthServer) SendEmailConfirmation(
	ctx context.Context,
	req *atomWebsite.SendEmailConfirmationRequest,
) (*atomWebsite.SendEmailConfirmationResponse, error) {
	err := r.sendEmailConfirmationUseCase.Execute(ctx, req.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "send failed: %s", err.Error())
	}

	return &atomWebsite.SendEmailConfirmationResponse{}, nil
}

func (r AuthServer) ConfirmEmail(
	ctx context.Context,
	req *atomWebsite.ConfirmEmailRequest,
) (*atomWebsite.ConfirmEmailResponse, error) {

	err := r.confirmEmailUseCase.Execute(
		ctx,
		usecase.ConfirmEmailRequest{
			UserEmail:        req.GetEmail(),
			ConfirmationCode: req.GetCode(),
		},
	)

	if err != nil {
		if errors.Is(err, usecase.ErrWrongCode) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &atomWebsite.ConfirmEmailResponse{}, nil
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
		Email:    req.GetEmail(),
		Password: req.Password,
		Phone:    req.GetPhone(),
	})

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Login fail: %s", err.Error())
	}

	return &atomWebsite.LoginResponse{
		Token: res.Token,
	}, nil
}

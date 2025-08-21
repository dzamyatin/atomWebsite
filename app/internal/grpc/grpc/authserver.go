package grpc

import (
	"context"
	atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"
	"github.com/dzamyatin/atomWebsite/internal/request"
	serviceauth "github.com/dzamyatin/atomWebsite/internal/service/auth"
	"github.com/dzamyatin/atomWebsite/internal/service/bus"
	"github.com/dzamyatin/atomWebsite/internal/transformer"
	"github.com/dzamyatin/atomWebsite/internal/usecase"
	"github.com/dzamyatin/atomWebsite/internal/validator"
	"github.com/guregu/null/v6"
	"github.com/pkg/errors"
)

type AuthServer struct {
	atomWebsite.UnimplementedAuthServer
	bus         bus.IBus
	validator   *validator.Validator
	transformer *transformer.Transformer

	registerUseCase              *usecase.Registration
	loginUseCase                 *usecase.Login
	confirmEmailUseCase          *usecase.ConfirmEmailUseCase
	sendEmailConfirmationUseCase *usecase.SendEmailConfirmationUseCase
	rememberPasswordUseCase      *usecase.RememberPasswordUseCase
	changePasswordUseCase        *usecase.ChangePasswordUseCase
	sendPhoneConfirmationUseCase *usecase.SendPhoneConfirmationUseCase
	confirmPhoneUseCase          *usecase.ConfirmPhoneUseCase
	auth                         serviceauth.IAuth
}

func NewAuthServer(
	registerUseCase *usecase.Registration,
	loginUseCase *usecase.Login,
	confirmEmailUseCase *usecase.ConfirmEmailUseCase,
	bus bus.IBus,
	sendEmailConfirmationUseCase *usecase.SendEmailConfirmationUseCase,
	rememberPasswordUseCase *usecase.RememberPasswordUseCase,
	validator *validator.Validator,
	changePasswordUseCase *usecase.ChangePasswordUseCase,
	sendPhoneConfirmationUseCase *usecase.SendPhoneConfirmationUseCase,
	confirmPhoneUseCase *usecase.ConfirmPhoneUseCase,
	transformer *transformer.Transformer,
	auth serviceauth.IAuth,
) AuthServer {
	return AuthServer{
		validator:                    validator,
		transformer:                  transformer,
		changePasswordUseCase:        changePasswordUseCase,
		registerUseCase:              registerUseCase,
		loginUseCase:                 loginUseCase,
		bus:                          bus,
		confirmEmailUseCase:          confirmEmailUseCase,
		sendEmailConfirmationUseCase: sendEmailConfirmationUseCase,
		rememberPasswordUseCase:      rememberPasswordUseCase,
		sendPhoneConfirmationUseCase: sendPhoneConfirmationUseCase,
		confirmPhoneUseCase:          confirmPhoneUseCase,
		auth:                         auth,
	}
}

func (r AuthServer) ConfirmPhone(
	ctx context.Context,
	req *atomWebsite.ConfirmPhoneRequest,
) (*atomWebsite.ConfirmPhoneResponse, error) {
	user, err := r.auth.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if err = r.validator.ValidateConfirmPhoneRequest(req); err != nil {
		return nil, r.ErrInvalidArgument(err)
	}

	err = r.confirmPhoneUseCase.Execute(ctx, r.transformer.TransformConfirmPhoneRequest(req, user))
	if err != nil {
		return nil, r.ErrInternal(err)
	}

	return &atomWebsite.ConfirmPhoneResponse{}, nil
}

func (r AuthServer) SendPhoneConfirmation(
	ctx context.Context,
	req *atomWebsite.SendPhoneConfirmationRequest,
) (*atomWebsite.SendPhoneConfirmationResponse, error) {
	if err := r.validator.ValidateSendPhoneConfirmationRequest(req); err != nil {
		return nil, r.ErrInvalidArgument(err)
	}

	err := r.sendPhoneConfirmationUseCase.Execute(ctx, r.transformer.TransformSendPhoneConfirmationRequest(req))
	if err != nil {
		return nil, r.ErrInternal(err)
	}

	return &atomWebsite.SendPhoneConfirmationResponse{}, nil
}

func (r AuthServer) ChangePassword(
	ctx context.Context,
	req *atomWebsite.ChangePasswordRequest,
) (*atomWebsite.ChangePasswordResponse, error) {
	if err := r.validator.ValidateChangePasswordRequest(req); err != nil {
		return nil, r.ErrInvalidArgument(err)
	}

	err := r.changePasswordUseCase.Execute(
		ctx,
		usecase.ChangePasswordRequest{
			Email:       req.GetEmail(),
			Phone:       req.GetPhone(),
			Code:        req.GetCode(),
			NewPassword: req.GetNewPassword(),
			OldPassword: req.GetOldPassword(),
		},
	)

	if err != nil {
		return nil, r.ErrInternal(err)
	}

	return &atomWebsite.ChangePasswordResponse{}, nil
}

func (r AuthServer) RememberPassword(
	ctx context.Context,
	req *atomWebsite.RememberPasswordRequest,
) (*atomWebsite.RememberPasswordResponse, error) {
	if err := r.validator.ValidateRememberPassword(req); err != nil {
		return nil, r.ErrInvalidArgument(err)
	}

	err := r.rememberPasswordUseCase.Execute(
		ctx,
		usecase.RememberPasswordRequest{
			Email: req.GetEmail(),
			Phone: req.GetPhone(),
		},
	)

	if err != nil {
		return nil, r.ErrInternal(err)
	}

	return &atomWebsite.RememberPasswordResponse{}, nil
}

func (r AuthServer) SendEmailConfirmation(
	ctx context.Context,
	req *atomWebsite.SendEmailConfirmationRequest,
) (*atomWebsite.SendEmailConfirmationResponse, error) {
	err := r.sendEmailConfirmationUseCase.Execute(ctx, req.GetEmail())
	if err != nil {
		return nil, r.ErrInternal(err)
	}

	return &atomWebsite.SendEmailConfirmationResponse{}, nil
}

func (r AuthServer) ConfirmEmail(
	ctx context.Context,
	req *atomWebsite.ConfirmEmailRequest,
) (*atomWebsite.ConfirmEmailResponse, error) {
	user, err := r.auth.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	err = r.confirmEmailUseCase.Execute(
		ctx,
		usecase.ConfirmEmailRequest{
			UserEmail:        req.GetEmail(),
			ConfirmationCode: req.GetCode(),
			CurrentUserUUID:  user.UUID,
		},
	)

	if err != nil {
		if errors.Is(err, usecase.ErrWrongCode) {
			return nil, r.ErrInvalidArgument(err)
		}

		return nil, r.ErrInternal(err)
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
		return nil, r.ErrInvalidArgument(err)
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
		return nil, r.ErrInvalidArgument(err)
	}

	return &atomWebsite.LoginResponse{
		Token: res.Token,
	}, nil
}

package usecase

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/request"
	"github.com/dzamyatin/atomWebsite/internal/response"
	serviceauth "github.com/dzamyatin/atomWebsite/internal/service/auth"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Login struct {
	logger   *zap.Logger
	provider serviceauth.IProvider
	jwt      serviceauth.IJWT
}

func NewLogin(
	logger *zap.Logger,
	provider serviceauth.IProvider,
	jwt serviceauth.IJWT,
) *Login {
	return &Login{logger: logger, provider: provider, jwt: jwt}
}

func (r *Login) Execute(ctx context.Context, req request.LoginRequest) (response.LoginResponse, error) {
	user, err := r.provider.GetUser(ctx, req)

	if err != nil {
		return response.LoginResponse{}, errors.Wrap(err, "provider get user")
	}

	jwt, err := r.jwt.CreateToken(*user)

	if err != nil {
		return response.LoginResponse{}, errors.Wrap(err, "create token")
	}

	return response.LoginResponse{
		Token: jwt.Value,
	}, nil
}

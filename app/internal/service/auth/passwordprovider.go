package serviceauth

import (
	"context"
	dtoauth "github.com/dzamyatin/atomWebsite/internal/dto/auth"
	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	"github.com/dzamyatin/atomWebsite/internal/request"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type PasswordProvider struct {
	rep                repository.IUserRepository
	passwordComparator entity.PasswordComparator
	logger             *zap.Logger
}

func NewPasswordProvider(rep repository.IUserRepository, logger *zap.Logger) *PasswordProvider {
	return &PasswordProvider{rep: rep, logger: logger}
}

func (r *PasswordProvider) GetUser(ctx context.Context, request request.LoginRequest) (*dtoauth.User, error) {
	u, err := r.rep.GetUserByEmail(ctx, request.Email)

	if err != nil {
		r.logger.Error("failed to get user by email", zap.Error(err))

		return nil, errors.Wrap(err, "failed to get user by email")
	}

	ok, err := u.CheckPassword(request.Password, r.passwordComparator)

	if err != nil {
		r.logger.Error("failed to check password", zap.Error(err))

		return nil, errors.Wrap(err, "failed to check password")
	}

	if ok {
		return dtoauth.NewUser(u.UUID), nil
	}

	return nil, ErrInvalidCredentials
}

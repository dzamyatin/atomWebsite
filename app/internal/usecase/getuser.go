package usecase

import (
	"context"

	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	"github.com/dzamyatin/atomWebsite/internal/request"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type GetUser struct {
	logger         *zap.Logger
	userRepository repository.IUserRepository
}

func NewGetUser(
	logger *zap.Logger,
	userRepository repository.IUserRepository,
) *GetUser {
	return &GetUser{
		logger:         logger,
		userRepository: userRepository,
	}
}

func (r *GetUser) Execute(ctx context.Context, req request.GetUserRequest) (request.GetUserResponse, error) {
	userUuid, err := uuid.Parse(req.UserUUID)
	if err != nil {
		return request.GetUserResponse{}, errors.Wrap(err, "pares uuid error")
	}

	user, err := r.userRepository.GetByUUID(ctx, entity.UserUuid(userUuid))
	if err != nil {
		return request.GetUserResponse{}, errors.Wrap(err, "get user by uuid error")
	}

	return request.GetUserResponse{
		Uuid:           uuid.UUID(user.UUID).String(),
		Email:          user.Email.V,
		Phone:          user.Phone.V,
		ConfirmedEmail: user.ConfirmedEmail,
		ConfirmedPhone: user.ConfirmedPhone,
	}, nil
}

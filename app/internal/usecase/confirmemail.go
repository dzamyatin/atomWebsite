package usecase

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	ErrWrongCode = errors.New("wrong code")
)

type ConfirmEmailRequest struct {
	UserEmail        string
	ConfirmationCode string
}

type ConfirmEmailUseCase struct {
	userRepository       repository.IUserRepository
	logger               *zap.Logger
	randomizerRepository repository.IRandomizerRepository
}

func NewConfirmEmailUseCase(
	userRepository repository.IUserRepository,
	logger *zap.Logger,
	randomizerRepository repository.IRandomizerRepository,
) *ConfirmEmailUseCase {
	return &ConfirmEmailUseCase{
		userRepository:       userRepository,
		logger:               logger,
		randomizerRepository: randomizerRepository,
	}
}

func (r *ConfirmEmailUseCase) Execute(ctx context.Context, req ConfirmEmailRequest) error {
	ok, err := r.randomizerRepository.CompareWithLast(ctx, req.UserEmail, req.ConfirmationCode)
	if err != nil {
		r.logger.Error("randomizer repository compare error", zap.Error(err))
		return errors.Wrap(err, "randomizer compare failed")
	}

	if !ok {
		return ErrWrongCode
	}

	user, err := r.userRepository.GetUserByEmail(ctx, req.UserEmail)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return ErrWrongCode
		}

		r.logger.Error("user repository get user by email error", zap.Error(err))
		return errors.Wrap(err, "user repository get user by email failed")
	}

	user.ConfirmedEmail = true

	err = r.userRepository.UpdateUser(ctx, *user)
	if err != nil {
		r.logger.Error("user repository update user error", zap.Error(err))
		return errors.Wrap(err, "user repository update user failed")
	}

	return nil
}

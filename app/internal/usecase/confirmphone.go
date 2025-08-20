package usecase

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	ErrUserNotFoundByPhone = errors.New("there is no user associated with this phone")
)

type ConfirmPhoneRequest struct {
	UserPhone        string
	ConfirmationCode string
}

type ConfirmPhoneUseCase struct {
	userRepository       repository.IUserRepository
	logger               *zap.Logger
	randomizerRepository repository.IRandomizerRepository
}

func NewConfirmPhoneUseCase(userRepository repository.IUserRepository, logger *zap.Logger, randomizerRepository repository.IRandomizerRepository) *ConfirmPhoneUseCase {
	return &ConfirmPhoneUseCase{userRepository: userRepository, logger: logger, randomizerRepository: randomizerRepository}
}

func (r *ConfirmPhoneUseCase) Execute(ctx context.Context, req ConfirmPhoneRequest) error {
	ok, err := r.randomizerRepository.CompareWithLast(ctx, req.UserPhone, req.ConfirmationCode)
	if err != nil {
		r.logger.Error("randomizer repository compare error", zap.Error(err))
		return errors.Wrap(err, "randomizer compare failed")
	}

	if !ok {
		return ErrWrongCode
	}

	user, err := r.userRepository.GetUserByPhone(ctx, req.UserPhone)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return ErrUserNotFoundByPhone
		}

		r.logger.Error("user repository get user by phone error", zap.Error(err))
		return errors.Wrap(err, "user repository get user by phone failed")
	}

	user.ConfirmedPhone = true

	err = r.userRepository.UpdateUser(ctx, *user)
	if err != nil {
		r.logger.Error("user repository update user error", zap.Error(err))
		return errors.Wrap(err, "user repository update user failed")
	}

	return nil
}

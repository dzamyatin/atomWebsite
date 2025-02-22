package usecase

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	ErrUserOrPasswordNotMatch = errors.New("user or password not match")
	ErrWrongConfirmationCode  = errors.New("wrong confirmation code")
)

type ChangePasswordRequest struct {
	Email       string
	Phone       string
	Code        string
	NewPassword string
	OldPassword string
}

type ChangePasswordUseCase struct {
	logger               *zap.Logger
	userRepository       repository.IUserRepository
	randomizerRepository repository.IRandomizerRepository
	passwordComparator   entity.PasswordComparator
	passwordEncoder      entity.PasswordEncoder
}

func NewChangePasswordUseCase(
	logger *zap.Logger,
	userRepository repository.IUserRepository,
	randomizerRepository repository.IRandomizerRepository,
	passwordComparator entity.PasswordComparator,
	passwordEncoder entity.PasswordEncoder,
) *ChangePasswordUseCase {
	return &ChangePasswordUseCase{
		logger:               logger,
		userRepository:       userRepository,
		randomizerRepository: randomizerRepository,
		passwordComparator:   passwordComparator,
		passwordEncoder:      passwordEncoder,
	}
}

func (r *ChangePasswordUseCase) Execute(ctx context.Context, req ChangePasswordRequest) error {

	if req.Email != "" {
		err := r.changePasswordByEmail(ctx, req)
		if err != nil {
			return errors.Wrap(err, "change password error")
		}

		return nil
	}

	if req.Phone != "" {
		err := r.changePasswordByPhone(ctx, req)
		if err != nil {
			return errors.Wrap(err, "change password error")
		}

		return nil
	}

	return errors.New("phone or email is required")
}

func (r *ChangePasswordUseCase) changePasswordByEmail(ctx context.Context, req ChangePasswordRequest) error {
	user, err := r.userRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return ErrUserOrPasswordNotMatch
		}
		r.logger.Error("get user by email", zap.Error(err))
		return errors.Wrap(err, "get user by email error")
	}

	//if user == nil {
	//	return ErrUserOrPasswordNotMatch
	//}

	if req.OldPassword != "" {
		ok, err := user.CheckPassword(req.OldPassword, r.passwordComparator)
		if err != nil {
			r.logger.Error("check password error", zap.Error(err))
			return errors.Wrap(err, "check password error")
		}

		if ok {
			return ErrUserOrPasswordNotMatch
		}

		err = r.changePassword(ctx, &user, req.NewPassword)
		if err != nil {
			r.logger.Error("change password error", zap.Error(err))
			return errors.Wrap(err, "change password error")
		}

		return nil
	}

	ok, err := r.randomizerRepository.CompareWithLast(ctx, req.Email, req.Code)
	if err != nil {
		r.logger.Error("randomizer compare error", zap.Error(err))
		return errors.Wrap(err, "randomizer compare error")
	}

	if !ok {
		return ErrWrongConfirmationCode
	}

	err = r.changePassword(ctx, &user, req.NewPassword)
	if err != nil {
		r.logger.Error("change password error", zap.Error(err))
		return errors.Wrap(err, "change password error")
	}

	return nil
}

func (r *ChangePasswordUseCase) changePassword(
	ctx context.Context,
	user *entity.User,
	pwd string,
) error {

	err := user.AddPassword(pwd, r.passwordEncoder)
	if err != nil {
		return errors.Wrap(err, "add password error")
	}

	err = r.userRepository.UpdateUser(ctx, *user)
	if err != nil {
		return errors.Wrap(err, "update user error")
	}

	return nil
}

func (r *ChangePasswordUseCase) changePasswordByPhone(ctx context.Context, req ChangePasswordRequest) error {
	user, err := r.userRepository.GetUserByPhone(ctx, req.Phone)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return ErrUserOrPasswordNotMatch
		}

		r.logger.Error("get user by phone error", zap.Error(err))
		return errors.Wrap(err, "get user by phone error")
	}

	//if user == nil {
	//	return ErrUserOrPasswordNotMatch
	//}

	if req.OldPassword != "" {
		ok, err := user.CheckPassword(req.OldPassword, r.passwordComparator)
		if err != nil {
			r.logger.Error("check password error", zap.Error(err))
			return errors.Wrap(err, "check password error")
		}

		if ok {
			return ErrUserOrPasswordNotMatch
		}

		err = r.changePassword(ctx, &user, req.NewPassword)
		if err != nil {
			r.logger.Error("change password error", zap.Error(err))
			return errors.Wrap(err, "change password error")
		}

		return nil
	}

	ok, err := r.randomizerRepository.CompareWithLast(ctx, req.Phone, req.Code)
	if err != nil {
		r.logger.Error("randomizer compare error", zap.Error(err))
		return errors.Wrap(err, "randomizer compare error")
	}

	if !ok {
		return ErrWrongConfirmationCode
	}

	err = r.changePassword(ctx, &user, req.NewPassword)
	if err != nil {
		r.logger.Error("change password error", zap.Error(err))
		return errors.Wrap(err, "change password error")
	}

	return nil
}

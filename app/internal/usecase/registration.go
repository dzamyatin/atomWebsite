package usecase

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/dto"
	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	"github.com/dzamyatin/atomWebsite/internal/validator"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

type RegistrationUseCase struct {
	userRepository  repository.IUserRepository
	passwordEncoder entity.PasswordEncoder
	validator       validator.IRegistrationValidator
	logger          *zap.Logger
}

func NewRegistrationUseCase(
	userRepository repository.IUserRepository,
	passwordEncoder entity.PasswordEncoder,
	validator validator.IRegistrationValidator,
	logger *zap.Logger,
) *RegistrationUseCase {
	return &RegistrationUseCase{
		userRepository:  userRepository,
		passwordEncoder: passwordEncoder,
		validator:       validator,
		logger:          logger,
	}
}

func (r *RegistrationUseCase) Execute(ctx context.Context, request dto.RegistrationRequest) error {
	if err := r.validate(ctx, request); err != nil {
		return err
	}

	user := entity.NewUserEntity(request.Email, request.Phone)

	if request.Password != "" {
		err := user.AddPassword(request.Password, r.passwordEncoder)
		if err != nil {
			return err
		}
	}

	err := r.userRepository.AddUser(ctx, *user)

	if err != nil {
		return err
	}

	return nil
}

func (r *RegistrationUseCase) validate(ctx context.Context, request dto.RegistrationRequest) error {
	if err := r.validator.Validate(request); err != nil {
		return err
	}

	_, err := r.userRepository.GetUserByPhone(ctx, request.Phone)

	if err == nil {
		return ErrUserAlreadyExists
	}

	if !errors.Is(err, repository.ErrUserNotFound) {
		r.logger.Error("User repository error get by phone", zap.Error(err))

		return errors.Wrap(err, "cant get user by phone")
	}

	_, err = r.userRepository.GetUserByEmail(ctx, request.Email)

	if err == nil {
		return ErrUserAlreadyExists
	}

	if !errors.Is(err, repository.ErrUserNotFound) {
		r.logger.Error("User repository error get by email", zap.Error(err))

		return errors.Wrap(err, "cant get user by email")
	}

	return nil
}

package usecase

import (
	"errors"
	"github.com/dzamyatin/atomWebsite/internal/dto"
	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	"github.com/dzamyatin/atomWebsite/internal/validator"
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

func (r *RegistrationUseCase) Execute(request dto.RegistrationRequest) error {
	if err := r.validate(request); err != nil {
		return err
	}

	user := entity.NewUserEntity(request.Email, request.Phone)

	if request.Password != "" {
		err := user.AddPassword(request.Password, r.passwordEncoder)
		if err != nil {
			return err
		}
	}

	err := r.userRepository.AddUser(user)

	if err != nil {
		return err
	}

	return nil
}

func (r *RegistrationUseCase) validate(request dto.RegistrationRequest) error {
	if err := r.validator.Validate(request); err != nil {
		return err
	}

	if _, err := r.userRepository.GetUserByPhone(request.Phone); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return ErrUserAlreadyExists
		}

		r.logger.Error("User repository error get by phone", zap.Error(err))

		return err
	}

	if _, err := r.userRepository.GetUserByEmail(request.Email); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return ErrUserAlreadyExists
		}

		r.logger.Error("User repository error get by email", zap.Error(err))

		return err
	}

	return nil
}

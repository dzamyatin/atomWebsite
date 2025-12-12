package usecase

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	"github.com/dzamyatin/atomWebsite/internal/request"
	servicemail "github.com/dzamyatin/atomWebsite/internal/service/mail"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

type Registration struct {
	userRepository  repository.IUserRepository
	passwordEncoder entity.PasswordEncoder
	//validator       validator.IRegistrationValidator
	logger                       *zap.Logger
	mailer                       servicemail.IMailer
	randomizerRepository         repository.IRandomizerRepository
	sendEmailConfirmationUseCase *SendEmailConfirmationUseCase
}

func NewRegistration(
	userRepository repository.IUserRepository,
	passwordEncoder entity.PasswordEncoder,
	//validator validator.IRegistrationValidator,
	logger *zap.Logger,
	mailer servicemail.IMailer,
	randomizerRepository repository.IRandomizerRepository,
	sendEmailConfirmationUseCase *SendEmailConfirmationUseCase,
) *Registration {
	return &Registration{
		userRepository:  userRepository,
		passwordEncoder: passwordEncoder,
		//validator:       validator,
		logger:                       logger,
		mailer:                       mailer,
		randomizerRepository:         randomizerRepository,
		sendEmailConfirmationUseCase: sendEmailConfirmationUseCase,
	}
}

func (r *Registration) Execute(ctx context.Context, request request.RegistrationRequest) error {
	if err := r.validate(ctx, request); err != nil {
		return err
	}

	user := entity.NewUser(
		request.Email,
		request.Phone,
	)

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

	if request.Email.V != "" {
		err = r.sendEmailConfirmationUseCase.Execute(ctx, request.Email.V)

		if err != nil {
			r.logger.Warn("failed to send email", zap.Error(err))
			return errors.Wrap(err, "send confirmation email")
		}
	}

	return nil
}

func (r *Registration) validate(ctx context.Context, request request.RegistrationRequest) error {
	//if err := r.validator.Validate(request); err != nil {
	//	return err
	//}

	if !request.Phone.Valid && !request.Email.Valid {
		return errors.New("one of phone or email should be specified")
	}

	if request.Phone.Valid {
		_, err := r.userRepository.GetUserByPhone(ctx, request.Phone.V)

		if err == nil {
			return ErrUserAlreadyExists
		}

		if !errors.Is(err, repository.ErrUserNotFound) {
			r.logger.Error("User repository error get by phone", zap.Error(err))

			return errors.Wrap(err, "cant get user by phone")
		}
	}

	if request.Email.Valid {
		_, err := r.userRepository.GetUserByEmail(ctx, request.Email.V)

		if err == nil {
			return ErrUserAlreadyExists
		}

		if !errors.Is(err, repository.ErrUserNotFound) {
			r.logger.Error("User repository error get by email", zap.Error(err))

			return errors.Wrap(err, "cant get user by email")
		}
	}

	return nil
}

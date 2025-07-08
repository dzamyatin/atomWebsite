package usecase

import (
	"context"
	"fmt"
	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	servicemail "github.com/dzamyatin/atomWebsite/internal/service/mail"
	servicemessenger "github.com/dzamyatin/atomWebsite/internal/service/messenger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"strings"
	"time"
)

type RememberPasswordRequest struct {
	Email string
	Phone string
}

type RememberPasswordUseCase struct {
	userRepository       repository.IUserRepository
	passwordEncoder      entity.PasswordEncoder
	logger               *zap.Logger
	mailer               servicemail.IMailer
	randomizerRepository repository.IRandomizerRepository
	messenger            servicemessenger.IMessengerService
}

func NewRememberPasswordUseCase(
	userRepository repository.IUserRepository,
	passwordEncoder entity.PasswordEncoder,
	logger *zap.Logger,
	mailer servicemail.IMailer,
	randomizerRepository repository.IRandomizerRepository,
	messenger servicemessenger.IMessengerService,
) *RememberPasswordUseCase {
	return &RememberPasswordUseCase{
		userRepository:       userRepository,
		passwordEncoder:      passwordEncoder,
		logger:               logger,
		mailer:               mailer,
		randomizerRepository: randomizerRepository,
		messenger:            messenger,
	}
}

func (r *RememberPasswordUseCase) Execute(ctx context.Context, req RememberPasswordRequest) error {
	group, ctx := errgroup.WithContext(ctx)

	if req.Email != "" {
		group.Go(func() error {
			return r.rememberByEmail(ctx, req.Email)
		})
	}

	if req.Phone != "" {
		group.Go(func() error {
			return r.rememberByPhone(ctx, req.Phone)
		})
	}

	return errors.Wrap(group.Wait(), "send remember code failed")
}

func (r *RememberPasswordUseCase) rememberByPhone(ctx context.Context, phone string) error {
	code, err := r.randomizerRepository.CreateRandomCode(ctx, phone, 1*time.Hour)
	if err != nil {
		return errors.Wrap(err, "create random code failed")
	}

	err = r.messenger.Send(ctx, phone, fmt.Sprintf("Your code is: %s", strings.ToUpper(code)))
	if err != nil {
		return errors.Wrap(err, "send remember code failed")
	}

	return nil
}

func (r *RememberPasswordUseCase) rememberByEmail(ctx context.Context, email string) error {
	code, err := r.randomizerRepository.CreateRandomCode(ctx, email, 1*time.Hour)
	if err != nil {
		return errors.Wrap(err, "create random code failed")
	}

	err = r.mailer.SendMail(
		ctx,
		email,
		"Remember password code",
		fmt.Sprintf("Your code is: %s", strings.ToUpper(code)),
	)

	if err != nil {
		r.logger.Error("send remember code failed", zap.Error(err))
		return errors.Wrap(err, "send remember code failed")
	}

	return nil
}

package usecase

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	servicemail "github.com/dzamyatin/atomWebsite/internal/service/mail"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strings"
	"time"
)

type SendEmailConfirmationUseCase struct {
	logger               *zap.Logger
	mailer               servicemail.IMailer
	randomizerRepository repository.IRandomizerRepository
}

func NewSendEmailConfirmationUseCase(
	logger *zap.Logger,
	mailer servicemail.IMailer,
	randomizerRepository repository.IRandomizerRepository,
) *SendEmailConfirmationUseCase {
	return &SendEmailConfirmationUseCase{logger: logger, mailer: mailer, randomizerRepository: randomizerRepository}
}

func (r *SendEmailConfirmationUseCase) Execute(
	ctx context.Context,
	email string,
) error {
	confirmationCode, err := r.randomizerRepository.CreateRandomCode(ctx, email, 1*time.Hour)
	if err != nil {
		r.logger.Warn("failed to generate confirmation code", zap.Error(err))
		return errors.Wrap(err, "failed to generate confirmation code")
	}

	err = r.mailer.SendMail(
		ctx,
		email,
		"Confirmation code",
		"Your confirmation code is "+strings.ToUpper(confirmationCode),
	)
	if err != nil {
		r.logger.Warn("failed to send email", zap.Error(err))
		return errors.Wrap(err, "send confirmation email")
	}

	return nil
}

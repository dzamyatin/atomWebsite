package usecase

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	servicemessengersender "github.com/dzamyatin/atomWebsite/internal/service/messenger/sender"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strings"
	"time"
)

type SendPhoneConfirmationUseCase struct {
	logger               *zap.Logger
	sender               servicemessengersender.ISenderService
	randomizerRepository repository.IRandomizerRepository
}

func NewSendPhoneConfirmationUseCase(logger *zap.Logger, sender servicemessengersender.ISenderService, randomizerRepository repository.IRandomizerRepository) *SendPhoneConfirmationUseCase {
	return &SendPhoneConfirmationUseCase{logger: logger, sender: sender, randomizerRepository: randomizerRepository}
}

func (r *SendPhoneConfirmationUseCase) Execute(
	ctx context.Context,
	phone string,
) error {
	count, err := r.randomizerRepository.CountCodes(ctx, phone)
	if err != nil {
		r.logger.Warn("failed to count codes for phone", zap.Error(err))
		return errors.Wrap(err, "failed to count codes for phone")
	}

	if count > maxSentConfirmation {
		return ErrTooManyConfirmations
	}

	confirmationCode, err := r.randomizerRepository.CreateRandomCode(ctx, phone, 1*time.Hour)
	if err != nil {
		r.logger.Warn("failed to generate confirmation code", zap.Error(err))
		return errors.Wrap(err, "failed to generate confirmation code")
	}

	err = r.sender.Send(ctx, phone, "Your confirmation code is "+strings.ToUpper(confirmationCode))

	if err != nil {
		r.logger.Warn("failed to send email", zap.Error(err))
		return errors.Wrap(err, "send confirmation email")
	}

	return nil
}

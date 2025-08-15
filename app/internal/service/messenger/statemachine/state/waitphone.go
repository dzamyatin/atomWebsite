package servicemessengerstatemachinestate

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	servicemessengermessage "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"
	"github.com/dzamyatin/atomWebsite/internal/service/messenger/statemachine"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type WaitForPhone struct {
	logger           *zap.Logger
	senderRepository repository.ISenderRepository
}

func NewWaitForPhone(
	logger *zap.Logger,
	senderRepository repository.ISenderRepository,
) *WaitForPhone {
	return &WaitForPhone{logger: logger, senderRepository: senderRepository}
}

func (r *WaitForPhone) State() servicemessengerstatemachine.StateName {
	return servicemessengerstatemachine.StateWaitPhone
}

func (r *WaitForPhone) ReceiveMessage(
	ctx context.Context,
	driver servicemessengermessage.IMessengerDriver,
	message servicemessengermessage.Message,
	machine servicemessengerstatemachine.IStateMachine,
) error {
	phone, err := driver.GetUserPhone(message)
	if err != nil {
		r.logger.Warn("Failed to get user phone", zap.Error(err))
		return errors.Wrap(err, "failed to get user phone")
	}

	if phone != "" {
		err = r.senderRepository.Save(
			ctx,
			entity.Sender{
				PhoneNumber: phone,
				Messenger:   driver.GetMessengerType(),
				Link:        message.ChatLink,
			},
		)
		if err != nil {
			r.logger.Warn("Failed to save user phone", zap.Error(err))
			return errors.Wrap(err, "failed to save user phone")
		}

		err = machine.Move(ctx, servicemessengerstatemachine.StatePhoneStored)
		if err != nil {
			r.logger.Warn("Failed to move user phone", zap.Error(err))
			return errors.Wrap(err, "failed to move state to phonestored")
		}

		if err = driver.SendMessage(
			servicemessengermessage.NewAnswer(
				message,
				"Your phone stored successfully",
			)); err != nil {
			r.logger.Error("failed to send message", zap.Error(err))
			return errors.Wrap(err, "driver.SendMessage")
		}

		return nil
	}

	if err := driver.SendMessage(
		servicemessengermessage.NewAnswer(
			message,
			"Please give me your phone!",
		)); err != nil {
		r.logger.Error("failed to send message", zap.Error(err))
		return errors.Wrap(err, "driver.SendMessage")
	}

	if err := driver.AskPhone(message.ChatLink); err != nil {
		r.logger.Error("failed to ask phone", zap.Error(err))
		return errors.Wrap(err, "driver.AskPhone")
	}

	return nil
}

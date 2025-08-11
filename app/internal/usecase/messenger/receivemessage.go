package usecasemessenger

import (
	"context"
	servicemessengerdriver "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"
	servicemessengerstatemachine "github.com/dzamyatin/atomWebsite/internal/service/messenger/statemachine"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ReceiveMessageInput struct {
	driver  servicemessengerdriver.IMessengerDriver
	message servicemessengerdriver.Message
}

func NewReceiveMessageInput(driver servicemessengerdriver.IMessengerDriver, message servicemessengerdriver.Message) ReceiveMessageInput {
	return ReceiveMessageInput{driver: driver, message: message}
}

type ReceiveMessageUseCase struct {
	logger       *zap.Logger
	stateFactory servicemessengerstatemachine.StateMachineFactory
}

func NewReceiveMessageUseCase(logger *zap.Logger, stateFactory servicemessengerstatemachine.StateMachineFactory) *ReceiveMessageUseCase {
	return &ReceiveMessageUseCase{logger: logger, stateFactory: stateFactory}
}

func (r *ReceiveMessageUseCase) Execute(ctx context.Context, input ReceiveMessageInput) error {
	chatID, err := input.driver.GetChatID(input.message)
	if err != nil {
		r.logger.Error("failed to get the chat ID", zap.Error(err))
		return errors.Wrap(err, "get chat id")
	}

	stateMachine, err := r.stateFactory.Load(
		ctx,
		string(input.message.MessengerType),
		chatID,
	)

	if err != nil {
		r.logger.Error("failed to load the chat ID", zap.Error(err))
		return errors.Wrap(err, "load the chat ID")
	}

	err = stateMachine.ReceiveMessage(ctx, input.driver, input.message, stateMachine)
	if err != nil {
		r.logger.Error("failed to receive the message", zap.Error(err))
		return errors.Wrap(err, "receive the message")
	}

	return nil
}

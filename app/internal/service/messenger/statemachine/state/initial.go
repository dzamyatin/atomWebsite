package servicemessengerstatemachinestate

import (
	"context"
	servicemessengermessage "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"
	"github.com/dzamyatin/atomWebsite/internal/service/messenger/statemachine"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type InitialState struct {
	logger *zap.Logger
}

func NewInitialState(logger *zap.Logger) *InitialState {
	return &InitialState{logger: logger}
}

func (r *InitialState) State() servicemessengerstatemachine.StateName {
	return servicemessengerstatemachine.StateInitial
}

func (r *InitialState) ReceiveMessage(
	ctx context.Context,
	driver servicemessengermessage.IMessengerDriver,
	message servicemessengermessage.Message,
	machine servicemessengerstatemachine.IStateMachine,
) error {
	err := driver.SendMessage(
		servicemessengermessage.NewAnswer(message, "Hello! What is your phone number?"),
	)
	if err != nil {
		r.logger.Warn("Failed to send message", zap.Error(err))
		return errors.Wrap(err, "failed to send message")
	}

	if err = machine.Move(ctx, servicemessengerstatemachine.StateWaitPhone); err != nil {
		r.logger.Error("move message", zap.Error(err))
		return errors.Wrap(err, "move message")
	}

	return nil
}

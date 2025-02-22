package servicemessengerstatemachinestate

import (
	"context"
	servicemessengermessage "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"
	"github.com/dzamyatin/atomWebsite/internal/service/messenger/statemachine"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type PhoneStoredState struct {
	logger *zap.Logger
}

func NewPhoneStoredState(logger *zap.Logger) *PhoneStoredState {
	return &PhoneStoredState{logger: logger}
}

func (r *PhoneStoredState) State() servicemessengerstatemachine.StateName {
	return servicemessengerstatemachine.StatePhoneStored
}

func (r *PhoneStoredState) ReceiveMessage(
	ctx context.Context,
	driver servicemessengermessage.IMessengerDriver,
	message servicemessengermessage.Message,
	machine servicemessengerstatemachine.IStateMachine,
) error {
	err := driver.SendMessage(
		servicemessengermessage.NewAnswer(message, "Your phone has been stored"),
	)
	if err != nil {
		r.logger.Warn("Failed to send message", zap.Error(err))
		return errors.Wrap(err, "failed to send message")
	}

	return nil
}

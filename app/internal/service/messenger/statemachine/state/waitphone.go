package servicemessengerstatemachinestate

import (
	"context"
	servicemessengermessage "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"
	"github.com/dzamyatin/atomWebsite/internal/service/messenger/statemachine"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type InitialAskPhone struct {
	logger *zap.Logger
	servicemessengerstatemachine.BaseState
	driver servicemessengermessage.IMessengerDriver
}

func (r *InitialAskPhone) State() servicemessengerstatemachine.StateName {
	return servicemessengerstatemachine.StateWaitPhone
}

func (r *InitialAskPhone) ReceiveMessage(
	ctx context.Context,
	message servicemessengermessage.Message,
	machine servicemessengerstatemachine.IStateMachine,
) error {
	if err := r.driver.SendMessage(
		servicemessengermessage.NewAnswer(
			message,
			"Please give me your phone!",
		)); err != nil {
		r.logger.Error("failed to send message", zap.Error(err))
		return errors.Wrap(err, "driver.SendMessage")
	}

	return nil
}

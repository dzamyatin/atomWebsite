package servicemessengerstatemachinestate

import (
	"context"
	servicemessengermessage "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"
	"github.com/dzamyatin/atomWebsite/internal/service/messenger/statemachine"
)

type InitialState struct {
	servicemessengerstatemachine.BaseState
}

func (r *InitialState) State() servicemessengerstatemachine.StateName {
	return servicemessengerstatemachine.StateInitial
}

func (r *InitialState) ReceiveMessage(
	ctx context.Context,
	message servicemessengermessage.IMessage,
	machine servicemessengerstatemachine.IStateMachine,
) error {
	return nil
}

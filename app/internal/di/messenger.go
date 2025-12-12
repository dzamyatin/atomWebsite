package di

import (
	servicemessengerstatemachine "github.com/dzamyatin/atomWebsite/internal/service/messenger/statemachine"
	servicemessengerstatemachinestate "github.com/dzamyatin/atomWebsite/internal/service/messenger/statemachine/state"
	"go.uber.org/zap"
)

func newStateRegistry(
	logger *zap.Logger,
	initialState *servicemessengerstatemachinestate.InitialState,
	waitForPhone *servicemessengerstatemachinestate.WaitForPhone,
	phoneStoredState *servicemessengerstatemachinestate.PhoneStoredState,
) servicemessengerstatemachine.IStateRegistry {
	return servicemessengerstatemachine.NewStateRegistry(
		logger,
		initialState,
		waitForPhone,
		phoneStoredState,
	)
}

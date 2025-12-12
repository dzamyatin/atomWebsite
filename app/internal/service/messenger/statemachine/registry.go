package servicemessengerstatemachine

import (
	"fmt"
	"go.uber.org/zap"
)

type IStateRegistry interface {
	Get(state StateName) (IState, error)
}

type StateRegistry struct {
	logger *zap.Logger
	states map[StateName]IState
}

func NewStateRegistry(logger *zap.Logger, states ...IState) *StateRegistry {
	stateMap := make(map[StateName]IState, len(states))

	for _, state := range states {
		stateMap[state.State()] = state
	}

	return &StateRegistry{logger: logger, states: stateMap}
}

func (r *StateRegistry) Get(state StateName) (IState, error) {
	v, ok := r.states[state]
	if !ok {
		r.logger.Error(fmt.Sprintf("state %s not found", state))
		return nil, fmt.Errorf("state %s not found", state)
	}

	return v, nil
}

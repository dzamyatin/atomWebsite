package servicemessengerstatemachine

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	servicemessengermessage "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type IStateMachine interface {
	Move(ctx context.Context, state StateName) error
}

type StateMachine struct {
	logger        *zap.Logger
	chat          entity.Chat
	currentState  IState
	chatRepo      repository.IChatRepository
	stateRegistry IStateRegistry
}

func NewStateMachine(
	logger *zap.Logger,
	chat entity.Chat,
	currentState IState,
	chatRepo repository.IChatRepository,
	stateRegistry IStateRegistry,
) *StateMachine {
	return &StateMachine{
		logger:        logger,
		chat:          chat,
		currentState:  currentState,
		chatRepo:      chatRepo,
		stateRegistry: stateRegistry,
	}
}

func (r *StateMachine) Move(ctx context.Context, state StateName) error {
	if r.currentState.State() == state {
		r.logger.Warn("try to move to the same state")
		return nil
	}

	newState, err := r.stateRegistry.Get(state)
	if err != nil {
		return errors.Wrap(err, "get state")
	}

	r.currentState = newState
	r.chat.State = string(state)

	err = r.chatRepo.Save(ctx, r.chat)
	if err != nil {
		return errors.Wrap(err, "save chat")
	}

	return nil
}

func (r *StateMachine) ReceiveMessage(
	ctx context.Context,
	message servicemessengermessage.Message,
	_ StateMachine,
) error {
	return r.currentState.ReceiveMessage(ctx, message, r)
}

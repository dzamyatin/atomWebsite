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
	Move(state StateName) error
	//Load(messenger, chatID string) (entity.Chat, error)
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

func (r *StateMachine) Load(messenger, chatID string) (entity.Chat, error) {
	chat, exist, err := r.chatRepo.Get(messenger, chatID)
	if err != nil {
		return entity.Chat{}, errors.Wrap(err, "load chat")
	}

	if !exist {
		chat = entity.Chat{
			Messenger: messenger,
			ChatID:    chatID,
			State:     string(StateInitial),
		}

		err = r.chatRepo.Save(chat)
		if err != nil {
			return entity.Chat{}, errors.Wrap(err, "save chat")
		}
	}

	return chat, nil
}

func (r *StateMachine) Move(state StateName) error {
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

	err = r.chatRepo.Save(r.chat)
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

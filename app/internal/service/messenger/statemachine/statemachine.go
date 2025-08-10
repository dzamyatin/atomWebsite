package servicemessengerstatemachine

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	servicemessengermessage "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"
	"github.com/pkg/errors"
)

type IStateMachine interface {
	Move(state StateName) error
}

type StateMachine struct {
	chat          entity.Chat
	currentState  IState
	chatRepo      repository.IChatRepository
	stateRegistry IStateRegistry
}

func (r *StateMachine) Load(messenger, phone, chatID string) (entity.Chat, error) {
	chat, exist, err := r.chatRepo.Get(messenger, phone)
	if err != nil {
		return entity.Chat{}, errors.Wrap(err, "load chat")
	}

	if !exist {
		chat = entity.Chat{
			Messenger: messenger,
			Phone:     phone,
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
	message servicemessengermessage.IMessage,
	_ StateMachine,
) error {
	return r.currentState.ReceiveMessage(ctx, message, r)
}

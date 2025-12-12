package servicemessengerstatemachine

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type StateMachineFactory struct {
	logger        *zap.Logger
	chatRepo      repository.IChatRepository
	stateRegistry IStateRegistry
}

func NewStateMachineFactory(
	logger *zap.Logger,
	chatRepo repository.IChatRepository,
	stateRegistry IStateRegistry,
) *StateMachineFactory {
	return &StateMachineFactory{logger: logger, chatRepo: chatRepo, stateRegistry: stateRegistry}
}

func (r *StateMachineFactory) Load(ctx context.Context, messenger, chatID string) (*StateMachine, error) {
	chat, exist, err := r.chatRepo.Get(ctx, messenger, chatID)
	if err != nil {
		return nil, errors.Wrap(err, "load chat")
	}

	if !exist {
		chat = entity.Chat{
			Messenger: messenger,
			ChatID:    chatID,
			State:     string(StateInitial),
		}

		err = r.chatRepo.Save(ctx, chat)
		if err != nil {
			r.logger.Error("save chat", zap.Error(err))
			return nil, errors.Wrap(err, "save chat")
		}
	}

	state, err := r.stateRegistry.Get(StateName(chat.State))
	if err != nil {
		r.logger.Error("load state fail", zap.Error(err))
		return nil, errors.Wrap(err, "get state")
	}

	return NewStateMachine(
		r.logger,
		chat,
		state,
		r.chatRepo,
		r.stateRegistry,
	), nil
}

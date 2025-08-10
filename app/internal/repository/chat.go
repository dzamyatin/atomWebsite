package repository

import (
	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/dzamyatin/atomWebsite/internal/service/db"
	"go.uber.org/zap"
)

type IChatRepository interface {
	Save(e entity.Chat) error
	Get(messenger, chatID string) (entity.Chat, bool, error)
}

type ChatRepository struct {
	logger *zap.Logger
	db     db.IDatabase
}

func (c ChatRepository) Save(e entity.Chat) error {
	//TODO implement me
	panic("implement me")
}

func (c ChatRepository) Get(messenger, chatID string) (entity.Chat, bool, error) {
	//TODO implement me
	panic("implement me")
}

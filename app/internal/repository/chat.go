package repository

import "github.com/dzamyatin/atomWebsite/internal/entity"

type IChatRepository interface {
	Save(e entity.Chat) error
	Get(messenger, chatID string) (entity.Chat, bool, error)
}

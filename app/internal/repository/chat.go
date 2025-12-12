package repository

import (
	"context"
	sql2 "database/sql"
	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/dzamyatin/atomWebsite/internal/service/db"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	ErrChatNotFound = errors.New("chat not found")
)

type IChatRepository interface {
	Save(ctx context.Context, e entity.Chat) error
	Get(ctx context.Context, messenger, chatID string) (entity.Chat, bool, error)
}

type ChatRepository struct {
	logger *zap.Logger
	db     db.IDatabase
}

func NewChatRepository(logger *zap.Logger, db db.IDatabase) *ChatRepository {
	return &ChatRepository{
		logger: logger,
		db:     db,
	}
}

func (r *ChatRepository) getSelectRows() []string {
	return r.getInsertRows()
}

func (r *ChatRepository) getInsertRows() []string {
	return []string{
		"messenger",
		"chat_id",
		"state",
	}
}

func (r *ChatRepository) getInsertVals(e entity.Chat) []any {
	return []any{
		e.Messenger,
		e.ChatID,
		e.State,
	}
}

func (r *ChatRepository) Save(ctx context.Context, e entity.Chat) error {
	ub := Builder.NewInsertBuilder()
	ub.InsertInto("chat")
	ub.Cols(r.getInsertRows()...)
	ub.Values(r.getInsertVals(e)...)
	buildOnConflictFields(
		ub,
		[]string{
			"messenger",
			"chat_id",
		},
		r.getInsertRows()...,
	)

	q, args := ub.Build()
	q = r.db.Rebind(q)

	_, err := r.db.Exec(ctx, q, args...)
	if err != nil {
		r.logger.Error("failed to update chat", zap.Error(err))
		return errors.Wrap(err, "update chat")
	}

	return nil
}

func (r *ChatRepository) Get(ctx context.Context, messenger, chatID string) (entity.Chat, bool, error) {
	sb := Builder.NewSelectBuilder()
	sb.From("chat")
	sb.Select(r.getSelectRows()...)
	sb.Where(
		sb.Equal("messenger", messenger),
		sb.Equal("chat_id", chatID),
	)
	sb.Limit(1)

	q, args := sb.Build()

	var chat entity.Chat
	err := r.db.Get(ctx, &chat, q, args...)
	if err != nil {
		if errors.Is(err, sql2.ErrNoRows) {
			return entity.Chat{}, false, nil
		}
		r.logger.Error("failed to get chat", zap.Error(err))
		return entity.Chat{}, false, errors.Wrap(err, "get chat")
	}

	return chat, true, nil
}

package repository

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/entity"
	"github.com/dzamyatin/atomWebsite/internal/service/db"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ISenderRepository interface {
	Save(ctx context.Context, sender entity.Sender) error
	GetByPhone(ctx context.Context, phone string) ([]entity.Sender, error)
}

type SenderRepository struct {
	logger *zap.Logger
	db     db.IDatabase
}

func NewSenderRepository(logger *zap.Logger, db db.IDatabase) *SenderRepository {
	return &SenderRepository{logger: logger, db: db}
}

func (r *SenderRepository) getVals(e entity.Sender) []any {
	return []any{
		e.PhoneNumber,
		e.Messenger,
		e.Link,
	}
}

func (r *SenderRepository) getCols() []string {
	return []string{
		"phone_number",
		"messenger",
		"link",
	}
}

func (r *SenderRepository) Save(
	ctx context.Context,
	sender entity.Sender,
) error {
	ib := Builder.NewInsertBuilder()
	ib.InsertInto("sender")

	ib.Cols(r.getCols()...)
	ib.Values(r.getVals(sender)...)

	buildOnConflictFields(
		ib,
		[]string{"phone_number", "messenger"},
		r.getCols()...,
	)

	q, args := ib.Build()

	if _, err := r.db.Exec(ctx, q, args...); err != nil {
		r.logger.Error("failed to insert sender", zap.Error(err))
		return errors.Wrap(err, "error saving sender")
	}

	return nil
}

func (r *SenderRepository) GetByPhone(ctx context.Context, phone string) ([]entity.Sender, error) {

	sb := Builder.NewSelectBuilder()
	sb.Select(r.getCols()...).From("sender")
	sb.Where(sb.Equal("phone_number", phone))

	q, args := sb.Build()

	var items []entity.Sender
	err := r.db.Select(ctx, &items, q, args...)
	if err != nil {
		r.logger.Error("failed to query sender", zap.Error(err))
		return nil, errors.Wrap(err, "error saving sender")
	}

	return items, nil
}

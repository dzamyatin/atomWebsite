package repository

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/service/db"
	"go.uber.org/zap"
	"time"
)

type ICounterRepository interface {
	Increase(
		ctx context.Context,
		key string,
	) (uint64, error)
}

type CounterRepository struct {
	logger *zap.Logger
	db     db.IDatabase
}

func NewCounterRepository(logger *zap.Logger, db db.IDatabase) *CounterRepository {
	return &CounterRepository{logger: logger, db: db}
}

func (r *CounterRepository) Increase(ctx context.Context, key string) (uint64, error) {
	ib := Builder.NewInsertBuilder()
	ib.InsertInto("counter")
	ib.Cols("key", "value", "created_at")
	ib.Values(key, 0, time.Now())
	ib.SQL(` ON CONFLICT (key) DO UPDATE SET value = value + 1`)
	ib.Returning("value")

	q, args := ib.Build()

	var count uint64
	err := r.db.Get(ctx, &count, q, args...)
	if err != nil {
		r.logger.Error("counter increase", zap.Error(err))
		return 0, err
	}

	return count, nil
}

package repository

import (
	"context"
	"database/sql"
	"github.com/dzamyatin/atomWebsite/internal/service/db"
	servicetime "github.com/dzamyatin/atomWebsite/internal/service/time"
	"github.com/pkg/errors"
	"math/rand/v2"
	"strings"

	"go.uber.org/zap"
	"time"
)

type IRandomizerRepository interface {
	CreateRandomCode(ctx context.Context, key string, ttl time.Duration) (string, error)
	CompareWithLast(ctx context.Context, key, code string) (bool, error)
}

type RandomizerRepository struct {
	logger *zap.Logger
	db     db.IDatabase
	time   servicetime.ITime
}

func NewRandomizerRepository(
	logger *zap.Logger,
	db db.IDatabase,
	time servicetime.ITime,
) *RandomizerRepository {
	return &RandomizerRepository{
		logger: logger,
		db:     db,
		time:   time,
	}
}

func (r *RandomizerRepository) CreateRandomCode(
	ctx context.Context,
	key string,
	ttl time.Duration,
) (string, error) {
	var code = make([]byte, 5)
	for i := range code {
		code[i] += rand.N[uint8](122-97) + 97
	}

	ib := Builder.NewInsertBuilder()
	ib.InsertInto("randomizer")
	ib.Cols("key", "code", "expired_at")
	ib.Values(strings.ToLower(key), code, time.Now().Add(ttl))

	q, args := ib.Build()

	_, err := r.db.Exec(ctx, q, args...)
	if err != nil {
		return "", err
	}

	return string(code), nil
}

func (r *RandomizerRepository) CompareWithLast(ctx context.Context, key, code string) (bool, error) {
	sb := Builder.NewSelectBuilder()
	sb.From("randomizer")
	sb.Select("key")
	sb.Where(
		sb.Equal("key", strings.ToLower(key)),
		sb.GT("expired_at", time.Now()),
		sb.Equal("code", strings.ToLower(code)),
	)
	sb.Limit(1)

	q, args := sb.Build()

	var res string
	err := r.db.Get(ctx, &res, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}

		r.logger.Error("couldn't execute query", zap.Error(err))
		return false, errors.Wrap(err, "check code code")
	}

	return true, nil
}

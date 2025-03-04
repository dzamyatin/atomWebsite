package usecasemigration

import (
	"context"
	"database/sql"
	"github.com/dzamyatin/atomWebsite/internal/migration"
	_ "github.com/dzamyatin/atomWebsite/internal/migration"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapgrpc"
)

type Up struct {
	logger *zap.Logger
	db     *sql.DB
}

func NewUp(
	logger *zap.Logger,
	db *sql.DB,
) *Up {
	return &Up{
		logger: logger,
		db:     db,
	}
}

func (r *Up) Execute(ctx context.Context) error {

	goose.SetBaseFS(migration.EmbedMigrations)

	goose.SetLogger(zapgrpc.NewLogger(r.logger))

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.UpContext(ctx, r.db, "sql"); err != nil {
		panic(err)
	}

	return nil
}

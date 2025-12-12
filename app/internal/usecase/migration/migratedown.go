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

type Down struct {
	logger *zap.Logger
	db     *sql.DB
}

func NewDown(
	logger *zap.Logger,
	db *sql.DB,
) *Down {
	return &Down{
		logger: logger,
		db:     db,
	}
}

func (r *Down) Execute(ctx context.Context) error {

	goose.SetBaseFS(migration.EmbedMigrations)

	goose.SetLogger(zapgrpc.NewLogger(r.logger))

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.DownContext(ctx, r.db, "sql"); err != nil {
		panic(err)
	}

	return nil
}

package migration

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(Up, Down)
}

func Up(_ context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("SELECT 1;")
	if err != nil {
		return err
	}
	return nil
}

func Down(_ context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("SELECT 1;")
	if err != nil {
		return err
	}
	return nil
}

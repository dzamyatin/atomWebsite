package db

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type IDatabase interface {
	Rebind(query string) string
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Begin(ctx context.Context) (*sqlx.Tx, error)
}

type Database struct {
	dbx *sqlx.DB
}

func NewDatabase(dbx *sqlx.DB) *Database {
	return &Database{dbx: dbx}
}

func (d *Database) Rebind(query string) string {
	return d.dbx.Rebind(query)
}

func (d *Database) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.dbx.SelectContext(ctx, dest, query, args...)
}

func (d *Database) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return d.dbx.GetContext(ctx, dest, query, args...)
}

func (d *Database) Query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	return d.dbx.QueryxContext(ctx, query, args...)
}

func (d *Database) QueryRow(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return d.dbx.QueryRowxContext(ctx, query, args...)
}

func (d *Database) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return d.dbx.ExecContext(ctx, query, args...)
}

func (d *Database) Begin(ctx context.Context) (*sqlx.Tx, error) {
	return d.dbx.BeginTxx(ctx, nil)
}

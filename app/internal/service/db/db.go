package db

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type idb interface {
	Rebind(query string) string
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Query(ctx context.Context, query string, args ...interface{}) (IRows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) IRow
	Exec(ctx context.Context, query string, args ...interface{}) (Result, error)
}

type IDatabase interface {
	idb
	Begin(ctx context.Context) (*Tx, error)
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

func (d *Database) Query(ctx context.Context, query string, args ...interface{}) (IRows, error) {
	rows, err := d.dbx.QueryxContext(ctx, query, args...)

	if err != nil {
		return nil, errors.Wrap(err, "failed to execute query")
	}

	return newRows(rows), nil
}

func (d *Database) QueryRow(ctx context.Context, query string, args ...interface{}) IRow {
	return newRow(d.dbx.QueryRowxContext(ctx, query, args...))
}

func (d *Database) Exec(ctx context.Context, query string, args ...interface{}) (Result, error) {
	res, err := d.dbx.ExecContext(ctx, query, args...)

	if err != nil {
		return Result{}, errors.Wrap(err, "failed to execute query")
	}

	return newResult(res), nil
}

func (d *Database) Begin(ctx context.Context) (*Tx, error) {
	tx, err := d.dbx.BeginTxx(ctx, nil)

	if err != nil {
		return nil, errors.Wrap(err, "failed to begin transaction")
	}

	return newTx(tx), nil
}

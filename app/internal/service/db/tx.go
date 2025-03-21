package db

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type ITx interface {
	idb
	Rollback() error
}

type Tx struct {
	*sqlx.Tx
}

func newTx(tx *sqlx.Tx) *Tx {
	return &Tx{
		Tx: tx,
	}
}

func (t *Tx) Exec(ctx context.Context, query string, args ...interface{}) (Result, error) {
	res, err := t.ExecContext(ctx, query, args...)

	if err != nil {
		return Result{}, err
	}

	return newResult(res), nil
}

func (t *Tx) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return t.GetContext(ctx, dest, query, args...)
}

func (t *Tx) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return t.SelectContext(ctx, dest, query, args...)
}

func (t *Tx) Query(ctx context.Context, query string, args ...interface{}) (IRows, error) {
	rows, err := t.Tx.QueryxContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	return newRows(rows), nil
}

func (t *Tx) QueryRow(ctx context.Context, query string, args ...interface{}) IRow {
	row := t.Tx.QueryRowxContext(ctx, query, args...)

	return newRow(row)
}

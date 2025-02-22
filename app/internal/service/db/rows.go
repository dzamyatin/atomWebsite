package db

import "github.com/jmoiron/sqlx"

type IRows interface {
	Close() error
	Next() bool
	StructScan(dest interface{}) error
}

type Rows struct {
	*sqlx.Rows
}

func newRows(rows *sqlx.Rows) *Rows {
	return &Rows{
		Rows: rows,
	}
}

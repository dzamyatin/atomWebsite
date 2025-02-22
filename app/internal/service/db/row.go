package db

import "github.com/jmoiron/sqlx"

type IRow interface {
	Scan(dest ...interface{}) error
}

type Row struct {
	*sqlx.Row
}

func newRow(row *sqlx.Row) *Row {
	return &Row{
		Row: row,
	}
}

package db

import (
	"context"

	servicetrace "github.com/dzamyatin/atomWebsite/internal/service/trace"
)

type TraceDatabaseDecorator struct {
	db    IDatabase
	trace *servicetrace.Trace
}

func NewTraceDatabaseDecorator(
	db IDatabase,
	trace *servicetrace.Trace,
) *TraceDatabaseDecorator {
	return &TraceDatabaseDecorator{
		db:    db,
		trace: trace,
	}
}

func (t *TraceDatabaseDecorator) Rebind(query string) string {
	return t.db.Rebind(query)
}

func (t *TraceDatabaseDecorator) Begin(ctx context.Context) (*Tx, error) {
	return t.db.Begin(ctx)
}

func (t *TraceDatabaseDecorator) Select(
	ctx context.Context,
	dest interface{},
	query string,
	args ...interface{},
) error {
	var res error
	t.trace.Trace(
		ctx,
		"db.Select",
		func(ctx context.Context) {
			res = t.db.Select(ctx, dest, query, args...)
		},
		servicetrace.NewTag("query", query),
	)

	return res
}

func (t *TraceDatabaseDecorator) Get(
	ctx context.Context,
	dest interface{},
	query string,
	args ...interface{},
) error {
	var res error
	t.trace.Trace(
		ctx,
		"db.Get",
		func(ctx context.Context) {
			res = t.db.Get(ctx, dest, query, args...)
		},
		servicetrace.NewTag("query", query),
	)

	return res
}

func (t *TraceDatabaseDecorator) Query(
	ctx context.Context,
	query string,
	args ...interface{},
) (rows IRows, err error) {
	t.trace.Trace(
		ctx,
		"db.Query",
		func(ctx context.Context) {
			rows, err = t.db.Query(ctx, query, args...)
		},
		servicetrace.NewTag("query", query),
	)

	return
}

func (t *TraceDatabaseDecorator) QueryRow(
	ctx context.Context,
	query string,
	args ...interface{},
) IRow {
	var res IRow
	t.trace.Trace(
		ctx,
		"db.QueryRow",
		func(ctx context.Context) {
			res = t.db.QueryRow(ctx, query, args...)
		},
		servicetrace.NewTag("query", query),
	)

	return res
}

func (t *TraceDatabaseDecorator) Exec(
	ctx context.Context,
	query string,
	args ...interface{},
) (res Result, err error) {
	t.trace.Trace(
		ctx,
		"",
		func(ctx context.Context) {
			res, err = t.db.Exec(ctx, query, args...)
		},
		servicetrace.NewTag("query", query),
	)

	return
}

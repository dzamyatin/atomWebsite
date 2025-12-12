package db

import (
	"database/sql"
)

type Result struct {
	sql.Result
}

func newResult(result sql.Result) Result {
	return Result{
		Result: result,
	}
}

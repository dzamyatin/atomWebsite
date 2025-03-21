package bus

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/service/db"
)

const BusPostgres BusName = "postgres"

type PostgresBus struct {
	MainBus
	db db.IDatabase
}

func (r *PostgresBus) Dispatch(ctx context.Context, command ICommand) error {
	panic("implement me")
}

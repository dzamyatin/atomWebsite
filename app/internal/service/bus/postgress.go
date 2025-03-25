package bus

import (
	"context"
	"encoding/json"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	"github.com/dzamyatin/atomWebsite/internal/service/db"
	"github.com/dzamyatin/atomWebsite/internal/service/metric"
	"github.com/huandu/go-sqlbuilder"
	"github.com/pkg/errors"
	"time"
)

const BusPostgres BusName = "postgres"

type BusStatus string

const (
	BusStatusNew        BusStatus = "new"
	BusStatusinProgress BusStatus = "in_progress"
	BusStatusSuccess    BusStatus = "success"
	BusStatusFailed     BusStatus = "failed"
)

type PostgresBus struct {
	queueName string
	BaseBus
	db db.IDatabase
}

func NewPostgresBus(queueName string, db db.IDatabase) *PostgresBus {
	return &PostgresBus{
		queueName: queueName,
		db:        db,
		BaseBus:   NewBaseBus(),
	}
}

func (r *PostgresBus) Dispatch(ctx context.Context, command ICommand) error {
	b, err := r.marshall(command)

	if err != nil {
		return errors.Wrap(err, "failed to marshal command")
	}

	ib := repository.Builder.NewInsertBuilder()
	ib.InsertInto("bus")
	ib.Cols(
		"uniqid",
		"queue",
		"created_at",
		"payload",
		"status",
	).
		Values(
			sqlbuilder.Raw("gen_random_uuid()"),
			r.queueName,
			time.Now(),
			b,
			BusStatusNew,
		)

	q, args := ib.Build()

	_, err = r.db.Exec(metric.MetricContext(ctx, "insert to bus"), q, args...)

	return errors.Wrap(err, "failed to dispatch command")
}

func (r *PostgresBus) marshall(command ICommand) ([]byte, error) {
	m, ok := command.(json.Marshaler)
	if ok {
		b, err := m.MarshalJSON()
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal command")
		}

		return b, nil
	}

	b, err := json.Marshal(command)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal command")
	}

	return b, nil
}

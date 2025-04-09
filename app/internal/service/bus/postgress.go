package bus

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	"github.com/dzamyatin/atomWebsite/internal/service/db"
	"github.com/dzamyatin/atomWebsite/internal/service/metric"
	"github.com/huandu/go-sqlbuilder"
	"github.com/pkg/errors"
	"go.uber.org/zap"
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

const (
	queueCheckDelay = time.Second * 1
)

var (
	ErrEmptyQueue = errors.New("empty queue")
)

type Item struct {
	Uniqid      string    `db:"uniqid"`
	Queue       string    `db:"queue"`
	CreatedAt   string    `db:"created_at"`
	Payload     string    `db:"payload"`
	Status      string    `db:"status"`
	TimeoutAt   time.Time `db:"timeout_at"`
	AttemptLeft uint64    `db:"attempt_left"`
	CommandName string    `db:"command_name"`
	RunAfter    string    `db:"run_after"`
}

type PostgresBus struct {
	queueName string
	BaseBus
	db        db.IDatabase
	timeout   time.Duration
	reattempt uint64
	logger    *zap.Logger
}

func NewPostgresBus(
	queueName string,
	db db.IDatabase,
	timeout time.Duration,
	reattempt uint64,
	logger *zap.Logger,
) *PostgresBus {
	return &PostgresBus{
		queueName: queueName,
		db:        db,
		BaseBus:   NewBaseBus(),
		timeout:   timeout,
		reattempt: reattempt,
		logger:    logger,
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
		"timeout_at",
		"attempt_left",
		"command_name",
		"run_after",
	).
		Values(
			sqlbuilder.Raw("gen_random_uuid()"),
			r.queueName,
			time.Now(),
			b,
			BusStatusNew,
			time.Now().Add(r.timeout),
			r.reattempt,
			command.GetName(),
			time.Now(),
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

func (r *PostgresBus) ProcessCycle(ctx context.Context, queueName string) error {
	for {
		select {
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "context canceled")
		default:
			err := r.process(ctx, queueName)

			if err != nil {
				if errors.Is(err, ErrEmptyQueue) {
					time.Sleep(queueCheckDelay) //use notify
					continue
				}
				r.logger.Error("failed to process command", zap.Error(err))
			}
		}
	}
}

func (r *PostgresBus) process(ctx context.Context, queueName string) error {
	item, err := r.extractItem(ctx, queueName)
	if err != nil {
		return errors.Wrap(err, "failed to extract item")
	}

	err = r.handleItem(ctx, item)
	if err != nil {
		err = r.failItem(ctx, item)
		if err != nil {
			return errors.Wrap(err, "failed to fail item")
		}

		return errors.Wrap(err, "failed to handle item")
	}

	err = r.successItem(ctx, item)
	if err != nil {
		return errors.Wrap(err, "failed to success item")
	}

	return nil
}

func (r *PostgresBus) handleItem(ctx context.Context, item *Item) error {
	command, err := r.GetCommand(item.CommandName)
	if err != nil {
		return errors.Wrap(err, "failed to get command")
	}

	err = json.Unmarshal([]byte(item.Payload), &command)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to unmarshal command %s", item.CommandName))
	}

	handler, err := r.GetHandler(command)
	if err != nil {
		return errors.Wrap(err, "failed to get handler")
	}

	return errors.Wrap(handler.Handle(ctx, command), "failed to handle command")
}

func (r *PostgresBus) successItem(ctx context.Context, item *Item) error {
	ub := repository.Builder.NewUpdateBuilder()
	ub.Set(
		ub.Assign("statue", BusStatusSuccess),
	)
	ub.Where(ub.E("uniqid", item.Uniqid))

	return nil
}

func (r *PostgresBus) extractItem(ctx context.Context, queueName string) (*Item, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to begin transaction")
	}
	defer tx.Rollback()

	item, ok, err := r.getItem(ctx, tx, queueName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get item")
	}

	if !ok {
		return nil, ErrEmptyQueue
	}

	err = r.startItem(ctx, tx, item)
	if err != nil {
		return nil, errors.Wrap(err, "failed to start item")
	}

	err = tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "failed to commit transaction")
	}

	if item.AttemptLeft < 0 {
		err = r.failItem(ctx, item)
		if err != nil {
			return nil, errors.Wrap(err, "failed to fail item")
		}

		return nil, errors.Wrap(err, "attempts exceeded")
	}

	return item, nil
}

func (r *PostgresBus) startItem(ctx context.Context, tx *db.Tx, item *Item) error {
	ub := repository.Builder.NewUpdateBuilder()
	item.AttemptLeft--
	ub.Update("bus").Set(
		ub.Assign("status", BusStatusinProgress),
		ub.Assign("queue", item.Queue),
		ub.Assign("timeout_at", time.Now().Add(r.timeout)),
		ub.Assign("attempt_left", item.AttemptLeft),
	)
	ub.Where(ub.E("uniqid", item.Uniqid))

	q, args := ub.Build()

	_, err := tx.Exec(metric.MetricContext(ctx, "start bus item"), q, args...)
	if err != nil {
		return errors.Wrap(err, "failed to start bus item")
	}

	return nil
}

func (r *PostgresBus) failItem(ctx context.Context, item *Item) error {
	ub := repository.Builder.NewUpdateBuilder()
	ub.Set(
		ub.Assign("statue", BusStatusFailed),
	)
	ub.Where(ub.E("uniqid", item.Uniqid))

	q, args := ub.Build()

	_, err := r.db.Exec(metric.MetricContext(ctx, "insert to bus"), q, args...)
	if err != nil {
		return errors.Wrap(err, "failed to fail item")
	}

	return nil
}

func (r *PostgresBus) getItem(ctx context.Context, tx *db.Tx, queueName string) (*Item, bool, error) {
	sb := repository.Builder.NewSelectBuilder()
	sb.Select(
		"uniqid",
		"queue",
		"created_at",
		"payload",
		"status",
		"timeout_at",
		"attempt_left",
		"command_name",
		"run_after",
	).
		From("bus").
		ForUpdate().
		Where(
			sb.LE("run_after", time.Now()),
			sb.GT("attempt_left", 0),
			sb.E("queue", queueName),
			sb.Or(
				sb.In("status", BusStatusNew, BusStatusFailed),
				sb.And(
					sb.E("status", BusStatusinProgress),
					sb.GTE("timeout_at", time.Now()),
				),
			),
		).
		Limit(1).
		OrderBy("created_at").
		Asc()

	q, args := sb.Build()

	q += " SKIP LOCKED"

	var item Item
	err := tx.GetContext(metric.MetricContext(ctx, "get item from bus"), &item, q, args...)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, false, errors.Wrap(err, "failed to execute query")
	}

	return &item, !errors.Is(err, sql.ErrNoRows), nil
}

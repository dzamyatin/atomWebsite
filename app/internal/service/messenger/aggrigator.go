package servicemessenger

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type AggrigatorMessenger struct {
	logger     *zap.Logger
	messengers []IMessengerService
}

func NewAggrigatorMessenger(
	logger *zap.Logger,
	messengers []IMessengerService,
) *AggrigatorMessenger {
	return &AggrigatorMessenger{
		logger:     logger,
		messengers: messengers,
	}
}

func (r *AggrigatorMessenger) Send(ctx context.Context, phone string, message string) error {
	var err error

	for _, m := range r.messengers {
		err = m.Send(ctx, phone, message)
		if err != nil {
			r.logger.Error("failed to send message", zap.Error(err))
		}
	}

	return errors.Wrap(err, "messenger send")
}

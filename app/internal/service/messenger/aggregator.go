package servicemessenger

import (
	"context"
	servicemessengermessage "github.com/dzamyatin/atomWebsite/internal/service/messenger/message"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type AggregatorSender struct {
	logger     *zap.Logger
	messengers []ISenderService
}

func NewAggrigatorSender(
	logger *zap.Logger,
	messengers []ISenderService,
) *AggregatorSender {
	return &AggregatorSender{
		logger:     logger,
		messengers: messengers,
	}
}

func (r *AggregatorSender) Send(ctx context.Context, phone string, message string) error {
	var err error

	for _, m := range r.messengers {
		err = m.Send(ctx, phone, message)
		if err != nil {
			r.logger.Error("failed to send message", zap.Error(err))
		}
	}

	return errors.Wrap(err, "messenger send")
}

func (r *AggregatorSender) Init(_ context.Context, data servicemessengermessage.IMessage) error {
	r.logger.Error("undefined sender to init", zap.Any("message", data))
	return nil
}

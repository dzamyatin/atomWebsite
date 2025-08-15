package servicemessengersender

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	servicemessengerdriver "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"
	messengerserver "github.com/dzamyatin/atomWebsite/internal/service/messenger/server"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type SequentiallySender struct {
	logger                  *zap.Logger
	counterRepository       repository.ICounterRepository
	senderRepository        repository.ISenderRepository
	messengerServerRegistry *messengerserver.MessengerServerRegistry
}

func NewSequentiallySender(
	logger *zap.Logger,
	counterRepository repository.ICounterRepository,
	senderRepository repository.ISenderRepository,
	messengerServerRegistry *messengerserver.MessengerServerRegistry,
) *SequentiallySender {
	return &SequentiallySender{
		logger:                  logger,
		counterRepository:       counterRepository,
		senderRepository:        senderRepository,
		messengerServerRegistry: messengerServerRegistry,
	}
}

func (r *SequentiallySender) Send(
	ctx context.Context,
	phone string,
	message string,
) error {
	senders, err := r.senderRepository.GetByPhone(ctx, phone)
	if err != nil {
		r.logger.Error("get sequentially sender by phone", zap.Error(err))
		return errors.Wrap(err, "get sequentially sender by phone")
	}

	count := len(senders)

	if count == 0 {
		return errors.New("senders not found")
	}

	counterValue, err := r.counterRepository.Increase(ctx, phone)

	chosen := counterValue % uint64(count)

	chosenSender := senders[chosen]

	sender, err := r.messengerServerRegistry.Get(chosenSender.Messenger)
	if err != nil {
		r.logger.Error("get sequentially sender by messenger", zap.Error(err))
		return errors.Wrap(err, "get sequentially sender by messenger")
	}

	err = sender.SendMessage(servicemessengerdriver.Message{
		ChatLink: chosenSender.Link,
		Text:     message,
	})

	if err != nil {
		r.logger.Error("send sequentially sender by messenger", zap.Error(err))
		return errors.Wrap(err, "send sequentially sender by messenger")
	}

	return nil
}

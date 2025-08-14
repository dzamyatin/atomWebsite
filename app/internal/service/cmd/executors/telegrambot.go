package executors

import (
	"context"
	mainarg "github.com/dzamyatin/atomWebsite/internal/service/arg"
	messengertelegram "github.com/dzamyatin/atomWebsite/internal/service/messenger/telegram"
	"github.com/dzamyatin/atomWebsite/internal/service/process"
	usecasemessenger "github.com/dzamyatin/atomWebsite/internal/usecase/messenger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ArgTelegramBotProcess struct {
	mainarg.Arg
}

type TelegramBotProcessCommand struct {
	logger            *zap.Logger
	telegramBotServer *messengertelegram.TelegramDriver
	receive           *usecasemessenger.ReceiveMessageUseCase
}

func NewTelegramBotProcessCommand(
	logger *zap.Logger,
	telegramBotServer *messengertelegram.TelegramDriver,
	receive *usecasemessenger.ReceiveMessageUseCase,
) *TelegramBotProcessCommand {
	return &TelegramBotProcessCommand{
		logger:            logger,
		telegramBotServer: telegramBotServer,
		receive:           receive,
	}
}

func (r *TelegramBotProcessCommand) Execute(ctx context.Context, u ArgTelegramBotProcess) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	return process.NewProcessManager(
		r.logger,
		process.Process{
			Name: "telegrambot-process",
			Object: process.NewProcessor(
				func(ctx context.Context) error {
					for update, err := range r.telegramBotServer.ReadMessages(ctx) {
						if err != nil {
							r.logger.Warn("telegrambot-process", zap.Error(err))
							return errors.Wrap(err, "bot process read messages")
						}

						err = r.receive.Execute(
							ctx,
							usecasemessenger.NewReceiveMessageInput(
								r.telegramBotServer,
								update,
							),
						)

						if err != nil {
							r.logger.Error("bot process execute", zap.Error(err))
							return errors.Wrap(err, "bot process execute")
						}
					}

					return nil
				},
				func() error {
					cancel()
					return nil
				},
			),
		},
	).Start(ctx)
}

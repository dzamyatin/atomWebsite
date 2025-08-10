package executors

import (
	"context"
	mainarg "github.com/dzamyatin/atomWebsite/internal/service/arg"
	messengertelegram "github.com/dzamyatin/atomWebsite/internal/service/messenger/telegram"
	"github.com/dzamyatin/atomWebsite/internal/service/process"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type ArgTelegramBotProcess struct {
	mainarg.Arg
}

type TelegramBotProcessCommand struct {
	logger            *zap.Logger
	telegramBotServer *messengertelegram.TelegramDriver
}

func NewTelegramBotProcessCommand(logger *zap.Logger, telegramBotServer *messengertelegram.TelegramDriver) *TelegramBotProcessCommand {
	return &TelegramBotProcessCommand{logger: logger, telegramBotServer: telegramBotServer}
}

func (r *TelegramBotProcessCommand) Execute(ctx context.Context, u ArgTelegramBotProcess) error {
	ctx, cancel := context.WithCancel(ctx)

	return process.NewProcessManager(
		r.logger,
		process.Process{
			Name: "telegrambot-process",
			Object: process.NewProcessor(
				func(ctx context.Context) error {
					return r.telegramBotServer.ReceiveUpdates(
						ctx,
						0,
						func(update tgbotapi.Update, bot *messengertelegram.TelegramDriver) error {
							return nil
						},
					)
				},
				func() error {
					cancel()
					return nil
				},
			),
		},
	).Start(ctx)
}

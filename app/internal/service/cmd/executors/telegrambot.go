package executors

import (
	"context"
	mainarg "github.com/dzamyatin/atomWebsite/internal/service/arg"
	messengertelegram "github.com/dzamyatin/atomWebsite/internal/service/messenger/telegram"
	"github.com/dzamyatin/atomWebsite/internal/service/process"
	"go.uber.org/zap"
)

type ArgTelegramBotProcess struct {
	mainarg.Arg
}

type TelegramBotProcessCommand struct {
	logger            *zap.Logger
	telegramBotServer *messengertelegram.TelegramBotServer
}

func NewTelegramBotProcessCommand(logger *zap.Logger, telegramBotServer *messengertelegram.TelegramBotServer) *TelegramBotProcessCommand {
	return &TelegramBotProcessCommand{logger: logger, telegramBotServer: telegramBotServer}
}

func (r *TelegramBotProcessCommand) Execute(ctx context.Context, u ArgTelegramBotProcess) error {
	return process.NewProcessManager(
		r.logger,
		process.Process{
			Name: "telegrambot-process",
			Object: process.NewProcessor(
				func(ctx context.Context) error {
					return r.telegramBotServer.Serve(ctx)
				},
				func() error {
					return r.telegramBotServer.Shutdown()
				},
			),
		},
	).Start(ctx)
}

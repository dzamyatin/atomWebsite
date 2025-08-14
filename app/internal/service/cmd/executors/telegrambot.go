package executors

import (
	"context"
	mainarg "github.com/dzamyatin/atomWebsite/internal/service/arg"
	servicemessengerdriver "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"
	messengertelegram "github.com/dzamyatin/atomWebsite/internal/service/messenger/telegram"
	"github.com/dzamyatin/atomWebsite/internal/service/process"
	usecasemessenger "github.com/dzamyatin/atomWebsite/internal/usecase/messenger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
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
							// add listening func to handle it in driver
							return r.receive.Execute(ctx, usecasemessenger.NewReceiveMessageInput(
								bot,
								servicemessengerdriver.Message{
									Username:      update.Message.From.UserName,
									MessengerType: servicemessengerdriver.MessengerTypeTelegram,
									ChatLink: servicemessengerdriver.ChatLink{
										Telegram: struct{ ChatID int64 }{
											ChatID: update.Message.Chat.ID,
										},
									},
									Text: update.Message.Text,
								},
							))

							//return nil
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

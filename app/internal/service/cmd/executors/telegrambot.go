package executors

import (
	"context"

	mainarg "github.com/dzamyatin/atomWebsite/internal/service/arg"
	servicemessengerdriver "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"
	messengerserver "github.com/dzamyatin/atomWebsite/internal/service/messenger/server"
	"github.com/dzamyatin/atomWebsite/internal/service/process"
	usecasemessenger "github.com/dzamyatin/atomWebsite/internal/usecase/messenger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ArgTelegramBotProcess struct {
	mainarg.Arg
	Bot string `arg:"-b,required" help:"name of driver"`
}

type TelegramBotProcessCommand struct {
	logger         *zap.Logger
	receive        *usecasemessenger.ReceiveMessageUseCase
	botRegistry    *messengerserver.MessengerServerRegistry
	processManager *process.ProcessShutdownerManager
}

func NewTelegramBotProcessCommand(
	logger *zap.Logger,
	botRegistry *messengerserver.MessengerServerRegistry,
	receive *usecasemessenger.ReceiveMessageUseCase,
	processManager *process.ProcessShutdownerManager,
) *TelegramBotProcessCommand {
	return &TelegramBotProcessCommand{
		logger:         logger,
		botRegistry:    botRegistry,
		receive:        receive,
		processManager: processManager,
	}
}

func (r *TelegramBotProcessCommand) Execute(ctx context.Context, u ArgTelegramBotProcess) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	bot, err := r.botRegistry.Get(servicemessengerdriver.MessengerType(u.Bot))
	if err != nil {
		return errors.Wrap(err, "get bot")
	}

	return r.processManager.Run(
		ctx,
		process.NewProcess(
			"bot-process",
			process.NewProcessor(
				func(ctx context.Context) error {
					for update, err := range bot.ReadMessages(ctx) {
						if err != nil {
							r.logger.Warn("bot-process", zap.Error(err))
							return errors.Wrap(err, "bot process read messages")
						}

						err = r.receive.Execute(
							ctx,
							usecasemessenger.NewReceiveMessageInput(
								bot,
								update,
							),
						)

						if err != nil {
							r.logger.Error("bot process execute", zap.Error(err))
							return errors.Wrap(err, "bot process execute")
						}
					}

					r.logger.Info("bot process execute finished")

					return nil
				},
				func() error {
					cancel()
					return nil
				},
			),
		),
	)
}

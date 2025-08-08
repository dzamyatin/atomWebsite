package servicemessenger

import (
	"context"
	servicemessengermessage "github.com/dzamyatin/atomWebsite/internal/service/messenger/message"
	"go.uber.org/zap"
)

type TelegramSender struct {
	logger *zap.Logger
}

func NewTelegramSender(
	logger *zap.Logger,
) *TelegramSender {
	return &TelegramSender{
		logger: logger,
	}
}

func (r *TelegramSender) Send(ctx context.Context, phone string, message string) error {
	return nil
}

func (r *TelegramSender) Init(ctx context.Context, data servicemessengermessage.IMessage) error {
	return nil
}

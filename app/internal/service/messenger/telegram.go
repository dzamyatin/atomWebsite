package servicemessenger

import (
	"context"
	"go.uber.org/zap"
)

type TelegramSender struct {
	logger *zap.Logger
}

func NewTelegramMessenger(
	logger *zap.Logger,
) *TelegramSender {
	return &TelegramSender{
		logger: logger,
	}
}

func (r *TelegramSender) Send(ctx context.Context, phone string, message string) error {
	return nil
}

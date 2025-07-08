package servicemessenger

import (
	"context"
	"go.uber.org/zap"
)

type TelegramMessenger struct {
	logger *zap.Logger
}

func NewTelegramMessenger(
	logger *zap.Logger,
) *TelegramMessenger {
	return &TelegramMessenger{
		logger: logger,
	}
}

func (r *TelegramMessenger) Send(ctx context.Context, phone string, message string) error {
	return nil
}

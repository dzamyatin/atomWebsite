package servicemessenger

import (
	"context"
	servicemessengermessage "github.com/dzamyatin/atomWebsite/internal/service/messenger/message"
)

type ISenderService interface {
	Send(ctx context.Context, phone string, message string) error
	Init(ctx context.Context, data servicemessengermessage.IMessage) error
}

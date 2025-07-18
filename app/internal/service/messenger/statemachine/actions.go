package servicemessengerstatemachine

import (
	"context"
	servicemessengermessage "github.com/dzamyatin/atomWebsite/internal/service/messenger/message"
)

type IStateActions interface {
	ReceiveMessage(ctx context.Context, message servicemessengermessage.IMessage) error
}

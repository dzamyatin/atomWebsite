package messengerserver

import (
	"context"
	servicemessengermessage "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"
	"iter"
)

type IMessengerServer interface {
	ReadMessages(ctx context.Context) iter.Seq2[servicemessengermessage.IMessage, error]
}

package messengerserver

import (
	"context"
	servicemessengermessage "github.com/dzamyatin/atomWebsite/internal/service/messenger/message"
	"iter"
)

type IMessengerServer interface {
	ReadMessages(ctx context.Context) iter.Seq2[servicemessengermessage.IMessage, error]
}

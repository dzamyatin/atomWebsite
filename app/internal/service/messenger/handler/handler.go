package servicemessengerhandler

import (
	"context"
	servicemessengermessage "github.com/dzamyatin/atomWebsite/internal/service/messenger/message"
)

type IServiceMessengerHandler interface {
	Handle(ctx context.Context, message servicemessengermessage.IMessage) error
}

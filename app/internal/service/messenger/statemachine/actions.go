package servicemessengerstatemachine

import (
	"context"
	servicemessengermessage "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"
)

type IStateActions interface {
	ReceiveMessage(
		ctx context.Context,
		message servicemessengermessage.IMessage,
		machine IStateMachine,
	) error
}

package servicemessengerstatemachine

import (
	"context"
	servicemessengermessage "github.com/dzamyatin/atomWebsite/internal/service/messenger/driver"
)

type IStateActions interface {
	ReceiveMessage(
		ctx context.Context,
		driver servicemessengermessage.IMessengerDriver,
		message servicemessengermessage.Message,
		machine IStateMachine,
	) error
}

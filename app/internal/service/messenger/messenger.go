package servicemessenger

import "context"

type IMessengerService interface {
	Send(ctx context.Context, phone string, message string) error
}

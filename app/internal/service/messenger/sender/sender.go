package servicemessengersender

import (
	"context"
)

type ISenderService interface {
	Send(ctx context.Context, phone string, message string) error
}

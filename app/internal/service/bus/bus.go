package bus

import (
	"context"
)

type BusName string

type IBus interface {
	Register(command ICommand, handler IHandler, name BusName) error
	Dispatch(ctx context.Context, command ICommand) error
}

type ICommand interface {
	GetName() string
}

type IHandler interface {
	Handle(ctx context.Context, command ICommand) error
}

package bus

import (
	"github.com/pkg/errors"
)

type BaseBus struct {
	dispatcher map[string]IHandler
}

func (r *BaseBus) Register(command ICommand, handler IHandler, name BusName) error {
	r.dispatcher[command.GetName()] = handler

	return nil
}

func (r *BaseBus) GetHandler(command ICommand) (IHandler, error) {
	handler, ok := r.dispatcher[command.GetName()]
	if !ok {
		return nil, errors.New("no Handler found")
	}

	return handler, nil
}

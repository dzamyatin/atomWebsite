package bus

import (
	"github.com/pkg/errors"
)

type BaseBus struct {
	dispatcher map[string]IHandler
	commands   map[string]ICommand
}

func NewBaseBus() BaseBus {
	return BaseBus{dispatcher: make(map[string]IHandler)}
}

func (r *BaseBus) Register(command ICommand, handler IHandler, name BusName) error {
	r.dispatcher[command.GetName()] = handler
	r.commands[command.GetName()] = command

	return nil
}

func (r *BaseBus) GetCommand(commandName string) (ICommand, error) {
	command, ok := r.commands[commandName]
	if !ok {
		return nil, errors.Errorf("command %s not found", commandName)
	}

	return command, nil
}

func (r *BaseBus) GetHandler(command ICommand) (IHandler, error) {
	handler, ok := r.dispatcher[command.GetName()]
	if !ok {
		return nil, errors.New("no Handler found")
	}

	return handler, nil
}

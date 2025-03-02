package command

import (
	"context"
	"fmt"
	"os"
)

type ICommand interface {
	Execute(ctx context.Context) int
	Help() string
}

var registry *Registry = NewRegistry()

func GetRegistry() *Registry {
	return registry
}

type Registry struct {
	commands       map[string]ICommand
	defaultCommand ICommand
}

func NewRegistry() *Registry {
	registry := &Registry{
		commands: make(map[string]ICommand),
	}

	return registry
}

func (r *Registry) Register(name string, command ICommand) {
	r.commands[name] = command
}

func (r *Registry) SetDefault(command ICommand) {
	r.defaultCommand = command
}

func (r *Registry) MustExecuteCommand(ctx context.Context, name string) {
	command, ok := r.commands[name]
	if !ok && r.defaultCommand != nil {
		r.defaultCommand.Execute(ctx)
		os.Exit(1)
	}
	if !ok {
		fmt.Printf("command not found %s", name)
		os.Exit(1)
	}

	os.Exit(command.Execute(ctx))
}

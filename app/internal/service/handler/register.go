package handler

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/service/bus"
	command2 "github.com/dzamyatin/atomWebsite/internal/service/command"
	"github.com/dzamyatin/atomWebsite/internal/usecase"
	"github.com/pkg/errors"
)

type RegisterHandler struct {
	registration *usecase.Registration
}

func NewRegisterHandler(registration *usecase.Registration) *RegisterHandler {
	return &RegisterHandler{registration: registration}
}

func (h *RegisterHandler) Handle(ctx context.Context, command bus.ICommand) error {
	v, ok := command.(*command2.RegisterCommand)
	if !ok {
		return errors.New("invalid command")
	}

	return errors.Wrap(h.registration.Execute(ctx, v.Req), "execute")
}

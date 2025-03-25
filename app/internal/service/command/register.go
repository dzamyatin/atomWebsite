package command

import (
	"context"
	"encoding/json"
	"github.com/dzamyatin/atomWebsite/internal/request"
	"github.com/dzamyatin/atomWebsite/internal/service/bus"
	"github.com/dzamyatin/atomWebsite/internal/usecase"
	"github.com/pkg/errors"
)

type RegisterCommand struct {
	Req request.RegistrationRequest `json:"req"`
}

func (c *RegisterCommand) GetName() string {
	return "RegisterCommand"
}

func (c *RegisterCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(*c)
}

func (c *RegisterCommand) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, c)
}

type RegisterHandler struct {
	registration *usecase.Registration
}

func NewRegisterHandler(registration *usecase.Registration) *RegisterHandler {
	return &RegisterHandler{registration: registration}
}

func (h *RegisterHandler) Handle(ctx context.Context, command bus.ICommand) error {
	v, ok := command.(*RegisterCommand)
	if !ok {
		return errors.New("invalid command")
	}

	return errors.Wrap(h.registration.Execute(ctx, v.Req), "execute")
}

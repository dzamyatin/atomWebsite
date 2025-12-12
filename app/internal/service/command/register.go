package command

import (
	"github.com/dzamyatin/atomWebsite/internal/request"
)

type RegisterCommand struct {
	Req request.RegistrationRequest `json:"req"`
}

func (c *RegisterCommand) GetName() string {
	return "RegisterCommand"
}

//func (c *RegisterCommand) MarshalJSON() ([]byte, error) {
//	return json.Marshal(*c)
//}
//
//func (c *RegisterCommand) UnmarshalJSON(b []byte) error {
//	return json.Unmarshal(b, c)
//}

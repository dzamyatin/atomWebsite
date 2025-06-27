package request

import "github.com/guregu/null/v6"

type RegistrationRequest struct {
	Email    null.Value[string] `json:"email"`
	Phone    null.Value[string] `json:"phone"`
	Password string             `json:"password"`
}

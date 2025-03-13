package validator

import (
	"errors"
	"github.com/dzamyatin/atomWebsite/internal/request"
)

var (
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPhone    = errors.New("invalid phone")
	ErrInvalidPassword = errors.New("invalid password")
)

type IRegistrationValidator interface {
	Validate(request request.RegistrationRequest) error
}

type RegistrationValidator struct{}

func NewRegistrationValidator() *RegistrationValidator {
	return &RegistrationValidator{}
}

func (r *RegistrationValidator) Validate(req request.RegistrationRequest) error {
	return nil
}

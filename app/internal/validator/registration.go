package validator

import (
	"errors"
	"github.com/dzamyatin/atomWebsite/internal/dto"
)

var (
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPhone    = errors.New("invalid phone")
	ErrInvalidPassword = errors.New("invalid password")
)

type IRegistrationValidator interface {
	Validate(request dto.RegistrationRequest) error
}

type RegistrationValidator struct{}

func NewRegistrationValidator() *RegistrationValidator {
	return &RegistrationValidator{}
}

func (r *RegistrationValidator) Validate(request dto.RegistrationRequest) error {
	return nil
}

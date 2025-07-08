package validator

import (
	"errors"
	atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"
	"regexp"
)

type Validator struct {
	emailRegex *regexp.Regexp
	phoneRegex *regexp.Regexp
}

func NewValidator() *Validator {
	emailRegex, err := regexp.Compile(
		`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`,
	)
	if err != nil {
		panic(err)
	}

	phoneRegex, err := regexp.Compile(`^\+[0-9]\d{10}$`)
	if err != nil {
		panic(err)
	}

	return &Validator{
		emailRegex: emailRegex,
		phoneRegex: phoneRegex,
	}
}

func (r Validator) ValidateRememberPassword(req *atomWebsite.RememberPasswordRequest) error {
	if req.GetEmail() == "" && req.GetPhone() == "" {
		return errors.New("email and phone are required")
	}

	if req.GetEmail() != "" {
		if !r.emailRegex.MatchString(req.GetEmail()) {
			return errors.New("invalid email")
		}
	}

	if req.GetPhone() != "" {
		if !r.phoneRegex.MatchString(req.GetPhone()) {
			return errors.New("invalid phone")
		}
	}

	return nil
}

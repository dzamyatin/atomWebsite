package validator

import (
	"errors"
	atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"
	"regexp"
)

type Validator struct {
	emailRegex    *regexp.Regexp
	phoneRegex    *regexp.Regexp
	passwordRegex *regexp.Regexp
}

func NewValidator() *Validator {
	passwordRegex, err := regexp.Compile(
		`^[A-Za-z\d!@#$%^&*()_+]{8,}$`,
	)
	if err != nil {
		panic(err)
	}

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
		emailRegex:    emailRegex,
		phoneRegex:    phoneRegex,
		passwordRegex: passwordRegex,
	}
}

func (r Validator) ValidateChangePasswordRequest(req *atomWebsite.ChangePasswordRequest) error {
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

	if req.GetOldPassword() == "" && req.GetCode() == "" {
		return errors.New("old password or verification code is required")
	}

	if req.GetNewPassword() == "" {
		return errors.New("new password is required")
	}

	if !r.passwordRegex.MatchString(req.GetNewPassword()) {
		return errors.New("new passwords should contain english letters, numbers, and special characters")
	}

	return nil
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

func (r Validator) ValidateConfirmPhoneRequest(req *atomWebsite.ConfirmPhoneRequest) error {
	if req.GetPhone() == "" {
		return errors.New("phone is required")
	}

	if !r.phoneRegex.MatchString(req.GetPhone()) {
		return errors.New("invalid phone")
	}

	if req.GetCode() == "" {
		return errors.New("code is required")
	}

	return nil
}

func (r Validator) ValidateSendPhoneConfirmationRequest(req *atomWebsite.SendPhoneConfirmationRequest) error {
	if req.GetPhone() == "" {
		return errors.New("phone is required")
	}

	if !r.phoneRegex.MatchString(req.GetPhone()) {
		return errors.New("invalid phone")
	}

	return nil
}

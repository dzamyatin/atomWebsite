package userservice

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type PasswordEncoder struct{}

func NewPasswordEncoder() *PasswordEncoder {
	return &PasswordEncoder{}
}

func (p *PasswordEncoder) Encode(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(b), err
}

func (p *PasswordEncoder) Compare(password string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return false, nil
	}

	return err == nil, err
}

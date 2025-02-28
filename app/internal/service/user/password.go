package userservice

import "golang.org/x/crypto/bcrypt"

type PasswordEncoder struct{}

func NewPasswordEncoder() *PasswordEncoder {
	return &PasswordEncoder{}
}

func (p *PasswordEncoder) Encode(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(b), err
}

func (p *PasswordEncoder) Compare(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

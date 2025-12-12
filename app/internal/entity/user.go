package entity

import (
	"github.com/google/uuid"
	"github.com/guregu/null/v6"
)

type UserUuid uuid.UUID

type User struct {
	UUID           UserUuid           `db:"uuid"`
	Email          null.Value[string] `db:"email"`
	Phone          null.Value[string] `db:"phone"`
	PasswordHash   string             `db:"password"`
	ConfirmedEmail bool               `db:"confirmed_email"`
	ConfirmedPhone bool               `db:"confirmed_phone"`
}

func NewUser(email, phone null.Value[string]) *User {
	return &User{
		Email: email,
		Phone: phone,
	}
}

type PasswordEncoder interface {
	Encode(password string) (string, error)
}

type PasswordComparator interface {
	Compare(password string, hash string) (ok bool, err error)
}

func (r *User) GenerateUUID() uuid.UUID {
	if r.UUID == [16]byte{} {
		r.UUID = UserUuid(uuid.New())
	}

	return uuid.UUID(r.UUID)
}

func (r *User) AddPassword(password string, passwordEncoder PasswordEncoder) error {
	encoded, err := passwordEncoder.Encode(password)

	if err != nil {
		return err
	}

	r.PasswordHash = encoded

	return nil
}

func (r *User) CheckPassword(password string, comparator PasswordComparator) (ok bool, err error) {
	return comparator.Compare(password, r.PasswordHash)
}

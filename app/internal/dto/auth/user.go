package dtoauth

import (
	"github.com/dzamyatin/atomWebsite/internal/entity"
)

type User struct {
	UUID entity.UserUuid
}

func NewUser(UUID entity.UserUuid) *User {
	return &User{UUID: UUID}
}

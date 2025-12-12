package dtoauth

import "github.com/google/uuid"

type User struct {
	UUID uuid.UUID
}

func NewUser(UUID uuid.UUID) *User {
	return &User{UUID: UUID}
}

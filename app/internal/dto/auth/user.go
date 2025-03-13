package dtoauth

type User struct {
	UUID string
}

func NewUser(UUID string) *User {
	return &User{UUID: UUID}
}

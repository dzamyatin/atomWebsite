package dtoauth

type Token struct {
	Value   string
	Payload *User
}

func NewToken(value string, payload *User) *Token {
	return &Token{Value: value, Payload: payload}
}

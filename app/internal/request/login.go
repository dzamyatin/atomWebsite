package request

type LoginRequest struct {
	Email    string
	Phone    string
	Password string
	Code     string
}

type LoginResponse struct {
	Token string
}

package request

type GetUserRequest struct {
	UserUUID string
}

type GetUserResponse struct {
	Uuid           string
	Email          string
	Phone          string
	ConfirmedEmail bool
	ConfirmedPhone bool
}

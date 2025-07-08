package usecase

type ChangePasswordRequest struct {
	Email       string
	Phone       string
	Code        string
	NewPassword string
	OldPassword string
}

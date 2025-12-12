package transformer

import (
	dtoauth "github.com/dzamyatin/atomWebsite/internal/dto/auth"
	atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"
	"github.com/dzamyatin/atomWebsite/internal/request"
	"github.com/dzamyatin/atomWebsite/internal/usecase"
)

type Transformer struct{}

func NewTransformer() *Transformer {
	return &Transformer{}
}

func (t *Transformer) TransformCurrentRequestFromUser(user dtoauth.User) request.GetUserRequest {
	return request.GetUserRequest{
		UserUUID: user.UUID.String(),
	}
}

func (t *Transformer) TransformGetUserResponse(res request.GetUserResponse) *atomWebsite.CurrentResponse {
	return &atomWebsite.CurrentResponse{
		Uuid:           res.Uuid,
		Email:          res.Email,
		Phone:          res.Phone,
		ConfirmedEmail: res.ConfirmedEmail,
		ConfirmedPhone: res.ConfirmedPhone,
	}
}

func (t *Transformer) TransformConfirmPhoneRequest(req *atomWebsite.ConfirmPhoneRequest, user dtoauth.User) usecase.ConfirmPhoneRequest {
	return usecase.ConfirmPhoneRequest{
		UserPhone:        req.GetPhone(),
		ConfirmationCode: req.GetCode(),
		CurrentUserUUID:  user.UUID,
	}
}

func (t *Transformer) TransformSendPhoneConfirmationRequest(req *atomWebsite.SendPhoneConfirmationRequest) string {
	return req.GetPhone()
}

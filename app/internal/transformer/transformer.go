package transformer

import (
	dtoauth "github.com/dzamyatin/atomWebsite/internal/dto/auth"
	atomWebsite "github.com/dzamyatin/atomWebsite/internal/grpc/generated"
	"github.com/dzamyatin/atomWebsite/internal/usecase"
)

type Transformer struct{}

func NewTransformer() *Transformer {
	return &Transformer{}
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

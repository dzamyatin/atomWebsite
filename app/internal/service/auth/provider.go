package serviceauth

import (
	"context"
	dtoauth "github.com/dzamyatin/atomWebsite/internal/dto/auth"
	"github.com/dzamyatin/atomWebsite/internal/request"
	"github.com/pkg/errors"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type IProvider interface {
	GetUser(ctx context.Context, request request.LoginRequest) (dtoauth.User, error)
}

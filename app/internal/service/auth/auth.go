package serviceauth

import (
	"context"
	dtoauth "github.com/dzamyatin/atomWebsite/internal/dto/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

var (
	ErrNotAuthorized = status.Error(codes.PermissionDenied, "not authorized")
)

type IAuth interface {
	GetUserFromCtx(ctx context.Context) (dtoauth.User, error)
}

type Auth struct {
	logger *zap.Logger
	jwt    IJWT
}

func NewAuth(logger *zap.Logger, jwt IJWT) *Auth {
	return &Auth{
		logger: logger,
		jwt:    jwt,
	}
}

func (r *Auth) GetUserFromCtx(ctx context.Context) (dtoauth.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return dtoauth.User{}, ErrNotAuthorized
	}

	authHeaders := md.Get("authorization")

	if len(authHeaders) == 0 {
		return dtoauth.User{}, ErrNotAuthorized
	}

	token, _ := strings.CutPrefix(authHeaders[0], "Bearer ")

	t, err := r.jwt.DecodeToken(token)
	if err != nil {
		r.logger.Error("failed to decode token", zap.Error(err))
		return dtoauth.User{}, ErrNotAuthorized
	}

	return *t.Payload, nil
}

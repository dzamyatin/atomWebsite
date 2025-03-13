package serviceauth

import (
	dtoauth "github.com/dzamyatin/atomWebsite/internal/dto/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

var (
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
)

const (
	claimUserUUID = "userUUID"
)

type IJWT interface {
	CreateToken(dtoauth.User) (*dtoauth.Token, error)
	DecodeToken(token string) (*dtoauth.Token, error)
}

type JWT struct {
	secret string
	jwtTTL time.Duration
	logger *zap.Logger
}

func (r *JWT) CreateToken(user dtoauth.User) (*dtoauth.Token, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, ErrUnexpectedSigningMethod
	}

	claims[claimUserUUID] = user.UUID

	t, err := token.SignedString([]byte(r.secret))

	if err != nil {
		return nil, errors.Wrap(err, "create token")
	}

	return &dtoauth.Token{
		Value:   t,
		Payload: &user,
	}, nil
}

func (r *JWT) DecodeToken(token string) (*dtoauth.Token, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigningMethod
		}

		return []byte(r.secret), nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "decode token")
	}

	claims, ok := t.Claims.(jwt.MapClaims)

	if !ok {
		return nil, ErrUnexpectedSigningMethod
	}

	uuid, ok := claims[claimUserUUID]

	if !ok {
		return nil, ErrUnexpectedSigningMethod
	}

	u, isString := uuid.(string)

	if !isString {
		return nil, ErrUnexpectedSigningMethod
	}

	return dtoauth.NewToken(
		token,
		dtoauth.NewUser(u),
	), nil
}

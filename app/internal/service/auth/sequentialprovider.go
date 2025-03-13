package serviceauth

import (
	"context"
	dtoauth "github.com/dzamyatin/atomWebsite/internal/dto/auth"
	"github.com/dzamyatin/atomWebsite/internal/request"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type SequentialProvider struct {
	logger    *zap.Logger
	providers []IProvider
}

func NewSequentialProvider(logger *zap.Logger, providers ...IProvider) *SequentialProvider {
	return &SequentialProvider{logger: logger, providers: providers}
}

func (s SequentialProvider) GetUser(ctx context.Context, request request.LoginRequest) (dtoauth.User, error) {
	for _, provider := range s.providers {
		user, err := provider.GetUser(ctx, request)
		if err != nil {
			if errors.Is(err, ErrInvalidCredentials) {
				continue
			}

			s.logger.Error("failed to get user", zap.Error(err))

			continue
		}

		return user, nil
	}

	return dtoauth.User{}, ErrInvalidCredentials
}

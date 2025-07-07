package usecase

import (
	"context"
	"github.com/dzamyatin/atomWebsite/internal/repository"
	"go.uber.org/zap"
)

type ConfirmEmailRequest struct {
	UserEmail        string
	ConfirmationCode string
}

type ConfirmEmailUseCase struct {
	userRepository       repository.IUserRepository
	logger               *zap.Logger
	randomizerRepository repository.IRandomizerRepository
}

func NewConfirmEmailUseCase(userRepository repository.IUserRepository, logger *zap.Logger, randomizerRepository repository.IRandomizerRepository) *ConfirmEmailUseCase {
	return &ConfirmEmailUseCase{userRepository: userRepository, logger: logger, randomizerRepository: randomizerRepository}
}

func (r ConfirmEmailUseCase) Execute(ctx context.Context, req ConfirmEmailRequest) error {

}

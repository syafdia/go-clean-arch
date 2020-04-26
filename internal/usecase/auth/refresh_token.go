package auth

import (
	"context"

	"github.com/syafdia/go-clean-arch/internal/entity"
	"github.com/syafdia/go-clean-arch/internal/repository"
	"github.com/syafdia/go-clean-arch/internal/usecase"
	"github.com/syafdia/go-clean-arch/pkg/log"
)

type refreshToken struct {
	tokenRepo repository.TokenRepository
}

func NewRefreshTokenUseCase(tokenRepo repository.TokenRepository) usecase.UseCase {
	return &refreshToken{tokenRepo}
}

func (s *refreshToken) Execute(ctx context.Context, input usecase.Input) (usecase.Output, error) {
	req, ok := input.(entity.RefreshTokenRequest)
	if !ok {
		return nil, entity.ErrFailedCastingUseCaseInput
	}

	log.Debug("request", req)
	// TODO

	return entity.RefreshTokenResponse{}, nil
}

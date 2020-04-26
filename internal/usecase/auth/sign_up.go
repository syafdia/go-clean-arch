package auth

import (
	"context"

	"github.com/syafdia/go-clean-arch/internal/entity"
	"github.com/syafdia/go-clean-arch/internal/repository"
	"github.com/syafdia/go-clean-arch/internal/usecase"
	"github.com/syafdia/go-clean-arch/pkg/log"
)

type signUp struct {
	userRepo  repository.UserRepository
	tokenRepo repository.TokenRepository
}

func NewSignUpUseCase(userRepo repository.UserRepository, tokenRepo repository.TokenRepository) usecase.UseCase {
	return &signUp{userRepo, tokenRepo}
}

func (s *signUp) Execute(ctx context.Context, input usecase.Input) (usecase.Output, error) {
	req, ok := input.(entity.SignUpRequest)
	if !ok {
		return nil, entity.ErrFailedCastingUseCaseInput
	}

	log.Debug("request", req)
	// TODO

	return entity.SignUpResponse{}, nil
}

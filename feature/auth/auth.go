package auth

import (
	repo "github.com/syafdia/go-clean-arch/data/repository"
	"github.com/syafdia/go-clean-arch/feature/auth/domain/usecase"
	"github.com/syafdia/go-clean-arch/feature/auth/presentation"
)

func NewAuthHTTP(userRepo repo.UserRepository, tokenRepo repo.TokenRepository) *presentation.AuthHTTP {
	return &presentation.AuthHTTP{
		AuthUseCase: usecase.NewAuthUseCase(userRepo, tokenRepo),
	}
}

package auth

import (
	repo "github.com/syafdia/go-clean-arch/data/repository"
	uc "github.com/syafdia/go-clean-arch/feature/auth/domain/usecase"
	"github.com/syafdia/go-clean-arch/feature/auth/presentation"
)

type AuthModule struct {
	userRepo repo.UserRepository
	authRepo repo.AuthRepository
}

func NewAuthModule(userRepo repo.UserRepository, authRepo repo.AuthRepository) *AuthModule {
	return &AuthModule{userRepo, authRepo}
}

func (a *AuthModule) authUseCase() uc.AuthUseCase {
	return uc.NewAuthUseCase(a.userRepo, a.authRepo)
}

func (a *AuthModule) AuthAPI() *presentation.AuthAPI {
	return presentation.NewAuthAPI(a.authUseCase())
}

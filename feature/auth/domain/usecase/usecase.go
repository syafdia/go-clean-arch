package usecase

import (
	repo "github.com/syafdia/go-clean-arch/data/repository"
	"github.com/syafdia/go-clean-arch/feature/auth/domain/model"
)

type AuthUseCase interface {
	SignIn(i model.InputSignIn) (model.ResponseSignIn, error)

	SignUp(i model.InputSignUp) error

	ForgotPassword() error
}

type authUseCase struct {
	userRepo repo.UserRepository
	authRepo repo.AuthRepository
}

func NewAuthUseCase(userRepo repo.UserRepository, authRepo repo.AuthRepository) *authUseCase {
	return &authUseCase{userRepo, authRepo}
}

func (a *authUseCase) SignIn(i model.InputSignIn) (model.ResponseSignIn, error) {
	u, err := a.authRepo.Authenticate(i.Username, i.Password)
	if err != nil {
		return model.EmptyResponseSignIn, err
	}

	t, err := a.authRepo.GenerateToken(u.Username, u.ID)
	if err != nil {
		return model.EmptyResponseSignIn, err
	}

	return model.ResponseSignIn{
		UserID:         u.ID,
		Username:       u.Username,
		Role:           u.Role,
		Token:          t.Value,
		TokenExpiredAt: t.ExpiredAt,
	}, err
}

func (a *authUseCase) SignUp(i model.InputSignUp) error {
	return nil
}

func (a *authUseCase) ForgotPassword() error {
	return nil
}

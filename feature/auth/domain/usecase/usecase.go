package usecase

import (
	repo "github.com/syafdia/go-clean-arch/data/repository"
	"github.com/syafdia/go-clean-arch/feature/auth/domain/model"
)

type AuthUseCase interface {
	SignIn(req model.SignInRequest) (model.SignInResponse, error)
	SignUp(req model.SignUpRequest) (model.SignUpResponse, error)
	RefreshToken(refreshToken string) (model.RefreshTokenResponse, error)
}

type authUseCase struct {
	userRepo  repo.UserRepository
	tokenRepo repo.TokenRepository
}

func NewAuthUseCase(userRepo repo.UserRepository, authRepo repo.TokenRepository) *authUseCase {
	return &authUseCase{userRepo, authRepo}
}

func (a *authUseCase) SignIn(i model.SignInRequest) (model.SignInResponse, error) {
	var resp model.SignInResponse

	user, err := a.userRepo.FindByUsernameAndPassword(i.Username, i.Password)
	if err != nil {
		return resp, err
	}

	token, err := a.tokenRepo.Generate(user.ID, user.Username)
	if err != nil {
		return resp, err
	}

	resp = model.SignInResponse{
		UserID:         user.ID,
		Username:       user.Username,
		Role:           user.Role,
		Token:          token.Value,
		RefreshToken:   token.RefreshToken,
		TokenExpiredAt: token.ExpiredAt,
	}

	return resp, nil
}

func (a *authUseCase) SignUp(i model.SignUpRequest) (model.SignUpResponse, error) {
	var resp model.SignUpResponse
	return resp, nil
}

func (a *authUseCase) RefreshToken(refreshToken string) (model.RefreshTokenResponse, error) {
	var resp model.RefreshTokenResponse
	return resp, nil
}

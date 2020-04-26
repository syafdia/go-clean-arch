package auth

import (
	"context"

	"github.com/syafdia/go-clean-arch/entity"
	"github.com/syafdia/go-clean-arch/repository"
	"github.com/syafdia/go-clean-arch/usecase"
)

type signIn struct {
	userRepo  repository.UserRepository
	tokenRepo repository.TokenRepository
}

func NewSignInUseCase(userRepo repository.UserRepository, tokenRepo repository.TokenRepository) usecase.UseCase {
	return &signIn{userRepo, tokenRepo}
}

func (s *signIn) Execute(ctx context.Context, input usecase.Input) (usecase.Output, error) {
	req, ok := input.(entity.SignInRequest)
	if !ok {
		return nil, entity.ErrFailedCastingUseCaseInput
	}

	var resp entity.SignInResponse

	user, err := s.userRepo.FindByUsernameAndPassword(ctx, req.Username, req.Password)
	if err != nil {
		return nil, entity.ErrUserNotFound
	}

	token, err := s.tokenRepo.Generate(ctx, user.ID, user.Username)
	if err != nil {
		return resp, entity.ErrFailedGeneratingToken
	}

	return entity.SignInResponse{
		UserID:         user.ID,
		Username:       user.Username,
		Role:           user.Role,
		Token:          token.Value,
		RefreshToken:   token.RefreshToken,
		TokenExpiredAt: token.ExpiredAt,
	}, nil
}

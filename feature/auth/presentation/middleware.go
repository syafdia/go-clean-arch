package presentation

import (
	"github.com/gin-gonic/gin"
	"github.com/syafdia/go-clean-arch/feature/auth/domain/usecase"
)

type AuthMiddleware struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthMiddleWare(authUseCase usecase.AuthUseCase) *AuthMiddleware {
	return &AuthMiddleware{authUseCase}
}

func (a *AuthMiddleware) ValidateOrDie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

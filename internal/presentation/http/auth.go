package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syafdia/go-clean-arch/internal/entity"
	"github.com/syafdia/go-clean-arch/internal/usecase"
)

type AuthHTTP interface {
	SignInAPIHandler(ctx *gin.Context)
	SignUpAPIHandler(ctx *gin.Context)
	TokenValidationMiddleware() gin.HandlerFunc
}

type authHTTP struct {
	signInUseCase       usecase.UseCase
	signUpUseCase       usecase.UseCase
	refreshTokenUseCase usecase.UseCase
}

func NewAuthHTTP(
	signInUseCase usecase.UseCase,
	signUpUseCase usecase.UseCase,
	refreshTokenUseCase usecase.UseCase,
) AuthHTTP {
	return &authHTTP{signInUseCase, signUpUseCase, refreshTokenUseCase}
}

func (a *authHTTP) SignInAPIHandler(ctx *gin.Context) {
	var req entity.SignInRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		WriteError(ctx, http.StatusBadRequest, err)
		return
	}

	resp, err := a.signInUseCase.Execute(ctx, req)
	if err != nil {
		WriteError(ctx, http.StatusBadRequest, err)
		return
	}

	WriteData(ctx, http.StatusOK, resp)
}

func (a *authHTTP) SignUpAPIHandler(ctx *gin.Context) {
	var req entity.SignUpRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		WriteError(ctx, http.StatusBadRequest, err)
		return
	}

	resp, err := a.signUpUseCase.Execute(ctx, req)
	if err != nil {
		WriteError(ctx, http.StatusBadRequest, err)
		return
	}

	WriteData(ctx, http.StatusOK, resp)
}

func (a *authHTTP) TokenValidationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

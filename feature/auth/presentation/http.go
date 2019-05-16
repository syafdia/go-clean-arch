package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syafdia/go-clean-arch/core/response"
	"github.com/syafdia/go-clean-arch/feature/auth/domain/model"
	"github.com/syafdia/go-clean-arch/feature/auth/domain/usecase"
)

type AuthHTTP struct {
	AuthUseCase usecase.AuthUseCase
}

func (a *AuthHTTP) SignInAPIHandler(ctx *gin.Context) {
	var req model.SignInRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.WriteError(ctx, http.StatusBadRequest, err)
		return
	}

	resp, err := a.AuthUseCase.SignIn(req)
	if err != nil {
		response.WriteError(ctx, http.StatusBadRequest, err)
		return
	}

	response.WriteData(ctx, http.StatusOK, resp)
}

func (a *AuthHTTP) SignUpAPIHandler(ctx *gin.Context) {
	var req model.SignUpRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.WriteError(ctx, http.StatusBadRequest, err)
		return
	}

	resp, err := a.AuthUseCase.SignUp(req)
	if err != nil {
		response.WriteError(ctx, http.StatusBadRequest, err)
		return
	}

	response.WriteData(ctx, http.StatusOK, resp)
}

func (a *AuthHTTP) TokenValidationMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}

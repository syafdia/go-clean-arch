package presentation

import (
	"net/http"

	"github.com/syafdia/go-clean-arch/core/logger"

	"github.com/gin-gonic/gin"
	"github.com/syafdia/go-clean-arch/core/response"
	"github.com/syafdia/go-clean-arch/feature/auth/domain/model"
	"github.com/syafdia/go-clean-arch/feature/auth/domain/usecase"
)

type AuthAPI struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthAPI(authUseCase usecase.AuthUseCase) *AuthAPI {
	return &AuthAPI{authUseCase}
}

func (a *AuthAPI) SignIn(ctx *gin.Context) {
	var input model.InputSignIn

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		logger.E(err)
		response.Write(ctx, http.StatusBadRequest, response.Response{
			Error: "username & password must be filled",
		})
		return
	}

	respSignIn, err := a.authUseCase.SignIn(input)
	if err != nil {
		logger.E(err)
		response.Write(ctx, http.StatusBadRequest, response.Response{
			Error: "invalid username or password",
		})
		return
	}

	response.Write(ctx, http.StatusBadRequest, response.Response{
		Data: respSignIn,
	})
}

func (a *AuthAPI) SignUp(ctx *gin.Context) {

}

func (a *AuthAPI) ForgotPassword(ctx *gin.Context) {

}

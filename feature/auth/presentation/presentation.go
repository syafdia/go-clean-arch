package presentation

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "username & password are required",
		})
		return
	}

	user, err := a.authUseCase.SignIn(input)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "user is not found",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (a *AuthAPI) SignUp(ctx *gin.Context) {

}

func (a *AuthAPI) ForgotPassword(ctx *gin.Context) {

}

type AuthMiddleware struct {
}

package main

import (
	"github.com/gin-gonic/gin"
	pauth "github.com/syafdia/go-clean-arch/feature/auth/presentation"
	"github.com/syafdia/go-clean-arch/module"
)

const (
	routeSignIn         = "/auth/signin"
	routeSignUp         = "/auth/signup"
	routeForgotPassword = "/auth/forgot_password"
)

func main() {
	appModule := module.NewAppModule()
	repoModule := module.NewRepositoryModule(appModule)
	useCaseModule := module.NewUseCaseModule(repoModule)
	presentationModule := module.NewPresentationModule(useCaseModule)

	r := gin.Default()

	setupPingApi(r)
	setupAuthApi(r, presentationModule.AuthAPI())
	setupAuthMidleware(r, presentationModule.AuthMiddleware())

	r.Run()
}

func setupPingApi(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func setupAuthApi(r *gin.Engine, authAPI *pauth.AuthAPI) {
	r.POST(routeSignIn, authAPI.SignIn)
	r.POST(routeSignUp, authAPI.SignUp)
	r.POST(routeForgotPassword, authAPI.ForgotPassword)
}

func setupAuthMidleware(r *gin.Engine, authMiddleware *pauth.AuthMiddleware) {
	r.Use(authMiddleware.ValidateOrDie())
}

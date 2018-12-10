package main

import (
	"github.com/gin-gonic/gin"
	"github.com/syafdia/go-clean-arch/feature/auth/presentation"
	"github.com/syafdia/go-clean-arch/module"
)

const (
	routeSignIn         = "/auth/signin"
	routeSignUp         = "/auth/signup"
	routeForgotPassword = "/auth/forgot_password"
)

func main() {
	appModule := module.NewAppModule()
	repoModule := createRepoModule(appModule)
	featureModule := createFeatureModule(repoModule)
	authModule := featureModule.AuthModule()

	r := gin.Default()

	setupPingApi(r)
	setupAuthApi(r, authModule.AuthAPI())

	r.Run()
}

func createRepoModule(appModule *module.AppModule) *module.RepositoryModule {
	return module.NewRepositoryModule(appModule.DB(), appModule.RedisClient())
}

func createFeatureModule(repoModule *module.RepositoryModule) *module.FeatureModule {
	return module.NewFeatureModule(repoModule)
}

func setupPingApi(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func setupAuthApi(r *gin.Engine, authAPI *presentation.AuthAPI) {
	r.POST(routeSignIn, authAPI.SignIn)
	r.POST(routeSignUp, authAPI.SignUp)
	r.POST(routeForgotPassword, authAPI.ForgotPassword)
}

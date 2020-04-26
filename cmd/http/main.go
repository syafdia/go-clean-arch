package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/syafdia/go-clean-arch/internal/config"
	"github.com/syafdia/go-clean-arch/internal/data"
	"github.com/syafdia/go-clean-arch/internal/presentation/http"
	"github.com/syafdia/go-clean-arch/internal/repository"
	"github.com/syafdia/go-clean-arch/internal/usecase"
	authusecase "github.com/syafdia/go-clean-arch/internal/usecase/auth"
)

var (
	sqlDB       *sql.DB
	redisClient *redis.Client
	mainCfg     config.MainConfig
)

var (
	userRepo  repository.UserRepository
	tokenRepo repository.TokenRepository
)

var (
	signInUseCase       usecase.UseCase
	signUpUseCase       usecase.UseCase
	refreshTokenUseCase usecase.UseCase
)

var (
	authHTTP http.AuthHTTP
)

func main() {
	initAppModule()
	initRepo()
	initUseCase()
	initHTTPPresentation()

	run()
}

func initAppModule() {
	sqlDB = &sql.DB{}
	redisClient = &redis.Client{}
	mainCfg = config.NewMainConfig()
}

func initRepo() {
	userRepo = data.NewUserRepository(sqlDB, redisClient)
	tokenRepo = data.NewtokenRepository(sqlDB, redisClient, mainCfg.Repository)
}

func initUseCase() {
	signInUseCase = authusecase.NewSignInUseCase(userRepo, tokenRepo)
	signUpUseCase = authusecase.NewSignUpUseCase(userRepo, tokenRepo)
	refreshTokenUseCase = authusecase.NewRefreshTokenUseCase(tokenRepo)
}

func initHTTPPresentation() {
	authHTTP = http.NewAuthHTTP(signInUseCase, signUpUseCase, refreshTokenUseCase)
}

func run() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/auth/sign_in", authHTTP.SignInAPIHandler)
	r.POST("/auth/sign_up", authHTTP.SignUpAPIHandler)
	r.Use(authHTTP.TokenValidationMiddleware())

	r.Run(":3000")
}

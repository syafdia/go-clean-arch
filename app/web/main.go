package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	repo "github.com/syafdia/go-clean-arch/data/repository"
	authfeature "github.com/syafdia/go-clean-arch/feature/auth"
)

var (
	sqlDB       *sql.DB
	redisClient *redis.Client
)

var (
	userRepo  repo.UserRepository
	tokenRepo repo.TokenRepository
)

func main() {
	initApp()
	initRepo()

	r := gin.Default()

	initAuthFeature(r)
	initPingPongFeature(r)

	r.Run()

}

func initApp() {
	// TODO
}

func initRepo() {
	userRepo = repo.NewUserRepository(sqlDB, redisClient)
	tokenRepo = repo.NewtokenRepository(sqlDB, redisClient)
}

func initAuthFeature(r *gin.Engine) {
	authHTTP := authfeature.NewAuthHTTP(userRepo, tokenRepo)

	r.POST("/auth/sign_in", authHTTP.SignInAPIHandler)
	r.POST("/auth/sign_up", authHTTP.SignUpAPIHandler)
	r.Use(authHTTP.TokenValidationMiddleware())
}

func initPingPongFeature(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

package repository

import (
	"database/sql"

	"github.com/go-redis/redis"

	"github.com/syafdia/go-clean-arch/data/entity"
)

type AuthRepository interface {
	RefreshToken(token string) (entity.Token, error)

	Authenticate(username string, password string) (entity.User, error)

	SignOut(token string) error
}

type authRepository struct {
	db          *sql.DB
	redisClient *redis.Client
}

func NewAuthRepository(db *sql.DB, redisClient *redis.Client) AuthRepository {
	return &authRepository{
		db:          db,
		redisClient: redisClient,
	}
}

func (a *authRepository) RefreshToken(token string) (entity.Token, error) {
	return entity.Token{}, nil
}

func (a *authRepository) Authenticate(username string, password string) (entity.User, error) {
	return entity.User{
		ID:       "12312",
		Username: username,
		Role:     entity.UserRoleDefault,
	}, nil
}

func (a *authRepository) SignOut(token string) error {
	return nil
}

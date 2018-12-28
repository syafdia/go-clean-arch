package repository

import (
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/syafdia/go-clean-arch/data/entity"
)

type InputCreateUser struct {
}

type InputUpdateUser struct {
}

type InputFindUser struct {
}

type UserRepository interface {
	Create(i InputCreateUser) (entity.User, error)

	Get(userID string) (entity.User, error)

	Delete(userID string) error
}

type userRepository struct {
	db          *sql.DB
	redisClient *redis.Client
}

func NewUserRepository(db *sql.DB, redisClient *redis.Client) *userRepository {
	return &userRepository{
		db:          db,
		redisClient: redisClient,
	}
}

func (u *userRepository) Create(i InputCreateUser) (entity.User, error) {
	user := entity.User{
		ID:       "1",
		Username: "Foo Bar",
		Role:     entity.UserRoleDefault,
	}

	return user, nil
}

func (u *userRepository) Delete(userID string) error {
	return nil
}

func (u *userRepository) Get(userID string) (entity.User, error) {
	user := entity.User{
		ID:       "1",
		Username: "Foo Bar",
		Role:     entity.UserRoleDefault,
	}

	return user, nil
}

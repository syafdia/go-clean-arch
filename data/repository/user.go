package repository

import (
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/syafdia/go-clean-arch/data/entity"
)

type UserRepository interface {
	FindByUsernameAndPassword(username string, password string) (entity.User, error)
	FindByID(userID string) (entity.User, error)
	FindByRole(userRole entity.UserRole) ([]entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
	Delete(userID string) error
}

type userRepository struct {
	db          *sql.DB
	redisClient *redis.Client
}

func NewUserRepository(db *sql.DB, redisClient *redis.Client) UserRepository {
	return &userRepository{
		db:          db,
		redisClient: redisClient,
	}
}

func (u *userRepository) FindByUsernameAndPassword(username string, password string) (entity.User, error) {
	user := entity.User{
		ID:       "1",
		Username: "Foo Bar",
		Role:     entity.UserRoleDefault,
	}

	return user, nil
}

func (u *userRepository) FindByID(userID string) (entity.User, error) {
	user := entity.User{
		ID:       "1",
		Username: "Foo Bar",
		Role:     entity.UserRoleDefault,
	}

	return user, nil
}

func (u *userRepository) FindByRole(userRole entity.UserRole) ([]entity.User, error) {
	users := []entity.User{
		entity.User{ID: "1", Username: "Foo Bar", Role: userRole},
		entity.User{ID: "2", Username: "Bar Baz", Role: userRole},
		entity.User{ID: "3", Username: "Baz Tar", Role: userRole},
		entity.User{ID: "4", Username: "Tar Qux", Role: userRole},
	}

	return users, nil
}

func (u *userRepository) Create(user *entity.User) error {
	return nil
}

func (u *userRepository) Update(user *entity.User) error {
	return nil
}

func (u *userRepository) Delete(userID string) error {
	return nil
}

package data

import (
	"context"
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/syafdia/go-clean-arch/internal/entity"
	"github.com/syafdia/go-clean-arch/internal/repository"
)

type userRepository struct {
	db          *sql.DB
	redisClient *redis.Client
}

func NewUserRepository(db *sql.DB, redisClient *redis.Client) repository.UserRepository {
	return &userRepository{
		db:          db,
		redisClient: redisClient,
	}
}

func (u *userRepository) FindByUsernameAndPassword(ctx context.Context, username string, password string) (entity.User, error) {
	user := entity.User{
		ID:       "1",
		Username: "Foo Bar",
		Role:     entity.UserRoleDefault,
	}

	return user, nil
}

func (u *userRepository) FindByID(ctx context.Context, userID string) (entity.User, error) {
	user := entity.User{
		ID:       "1",
		Username: "Foo Bar",
		Role:     entity.UserRoleDefault,
	}

	return user, nil
}

func (u *userRepository) FindByRole(ctx context.Context, userRole entity.UserRole) ([]entity.User, error) {
	users := []entity.User{
		entity.User{ID: "1", Username: "Foo Bar", Role: userRole},
		entity.User{ID: "2", Username: "Bar Baz", Role: userRole},
		entity.User{ID: "3", Username: "Baz Tar", Role: userRole},
		entity.User{ID: "4", Username: "Tar Qux", Role: userRole},
	}

	return users, nil
}

func (u *userRepository) Create(ctx context.Context, user entity.User) error {
	return nil
}

func (u *userRepository) Update(ctx context.Context, user entity.User) error {
	return nil
}

func (u *userRepository) Delete(ctx context.Context, userID string) error {
	return nil
}

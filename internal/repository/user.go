package repository

import (
	"context"

	"github.com/syafdia/go-clean-arch/internal/entity"
)

type UserRepository interface {
	FindByUsernameAndPassword(ctx context.Context, username string, password string) (entity.User, error)
	FindByID(ctx context.Context, userID string) (entity.User, error)
	FindByRole(ctx context.Context, userRole entity.UserRole) ([]entity.User, error)
	Create(ctx context.Context, user entity.User) error
	Update(ctx context.Context, user entity.User) error
	Delete(ctx context.Context, userID string) error
}

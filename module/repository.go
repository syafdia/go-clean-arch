package module

import (
	"database/sql"

	"github.com/go-redis/redis"
	repo "github.com/syafdia/go-clean-arch/data/repository"
)

type RepositoryModule struct {
	db          *sql.DB
	redisClient *redis.Client
}

func NewRepositoryModule(db *sql.DB, redisClient *redis.Client) *RepositoryModule {
	return &RepositoryModule{
		db:          db,
		redisClient: redisClient,
	}
}

func (r *RepositoryModule) UserRepository() repo.UserRepository {
	userRepo := repo.NewUserRepository(r.db, r.redisClient)

	return userRepo
}

func (r *RepositoryModule) AuthRepository() repo.AuthRepository {
	authRepo := repo.NewAuthRepository(r.db, r.redisClient)

	return authRepo
}

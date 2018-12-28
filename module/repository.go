package module

import (
	"sync"

	repo "github.com/syafdia/go-clean-arch/data/repository"
)

type RepositoryModule struct{}

var (
	repoModuleInstance *RepositoryModule
	repoModuleOnce     sync.Once
	userRepo           repo.UserRepository
	authRepo           repo.AuthRepository
)

func NewRepositoryModule(appModule *AppModule) *RepositoryModule {
	repoModuleOnce.Do(func() {
		repoModuleInstance = &RepositoryModule{}
		userRepo = repo.NewUserRepository(appModule.DB(), appModule.RedisClient())
		authRepo = repo.NewAuthRepository(appModule.DB(), appModule.RedisClient())
	})

	return repoModuleInstance
}

func (r *RepositoryModule) UserRepository() repo.UserRepository {
	return userRepo
}

func (r *RepositoryModule) AuthRepository() repo.AuthRepository {
	return authRepo
}

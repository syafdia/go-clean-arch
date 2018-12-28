package module

import (
	"sync"

	ucAuth "github.com/syafdia/go-clean-arch/feature/auth/domain/usecase"
)

type UseCaseModule struct{}

var (
	useCaseModuleInstance *UseCaseModule
	useCaseModuleOnce     sync.Once
	authUseCaseInstance   ucAuth.AuthUseCase
)

func NewUseCaseModule(repoModule *RepositoryModule) *UseCaseModule {
	useCaseModuleOnce.Do(func() {
		useCaseModuleInstance = &UseCaseModule{}
		authUseCaseInstance = ucAuth.NewAuthUseCase(
			repoModule.UserRepository(),
			repoModule.AuthRepository(),
		)
	})

	return useCaseModuleInstance
}

func (u *UseCaseModule) AuthUseCase() ucAuth.AuthUseCase {
	return authUseCaseInstance
}

package module

import (
	"sync"

	presAuth "github.com/syafdia/go-clean-arch/feature/auth/presentation"
)

type PresentationModule struct{}

var (
	presentationModuleInstance *PresentationModule
	presentationModuleOnce     sync.Once
	authAPIInstance            *presAuth.AuthAPI
	authMiddlewareInstance     *presAuth.AuthMiddleware
)

func NewPresentationModule(useCaseModule *UseCaseModule) *PresentationModule {
	presentationModuleOnce.Do(func() {
		presentationModuleInstance = &PresentationModule{}
		authAPIInstance = presAuth.NewAuthAPI(useCaseModule.AuthUseCase())
		authMiddlewareInstance = presAuth.NewAuthMiddleWare(useCaseModule.AuthUseCase())
	})

	return presentationModuleInstance
}

func (p *PresentationModule) AuthAPI() *presAuth.AuthAPI {
	return authAPIInstance
}

func (p *PresentationModule) AuthMiddleware() *presAuth.AuthMiddleware {
	return authMiddlewareInstance
}

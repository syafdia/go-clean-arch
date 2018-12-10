package module

import "github.com/syafdia/go-clean-arch/feature/auth"

type FeatureModule struct {
	repoModule *RepositoryModule
}

func NewFeatureModule(repoModule *RepositoryModule) *FeatureModule {
	return &FeatureModule{repoModule}
}

func (f *FeatureModule) AuthModule() *auth.AuthModule {
	return auth.NewAuthModule(f.repoModule.UserRepository(), f.repoModule.AuthRepository())
}

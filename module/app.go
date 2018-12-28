package module

import (
	"database/sql"
	"sync"

	"github.com/go-redis/redis"
)

type AppModule struct{}

var (
	appModuleInstance *AppModule
	appModuleOnce     sync.Once
	sqlDB             *sql.DB
	redisClient       *redis.Client
)

func NewAppModule() *AppModule {
	appModuleOnce.Do(func() {
		appModuleInstance = &AppModule{}
	})

	return appModuleInstance
}

func (a *AppModule) DB() *sql.DB {
	return nil
}

func (a *AppModule) RedisClient() *redis.Client {
	return nil
}

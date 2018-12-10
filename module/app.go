package module

import (
	"database/sql"

	"github.com/go-redis/redis"
)

type AppModule struct{}

func NewAppModule() *AppModule {
	return &AppModule{}
}

func (a *AppModule) DB() *sql.DB {
	return nil
}

func (a *AppModule) RedisClient() *redis.Client {
	return nil
}

package data

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-redis/redis"
	"github.com/syafdia/go-clean-arch/internal/config"
	"github.com/syafdia/go-clean-arch/internal/entity"
	"github.com/syafdia/go-clean-arch/internal/repository"
)

type tokenRepository struct {
	db          *sql.DB
	redisClient *redis.Client
	config      config.RepositoryConfig
}

func NewtokenRepository(db *sql.DB, redisClient *redis.Client, config config.RepositoryConfig) repository.TokenRepository {
	return &tokenRepository{db, redisClient, config}
}

func (a *tokenRepository) Generate(ctx context.Context, userID string, username string) (entity.Token, error) {
	return entity.Token{
		UserID:       userID,
		Value:        "thisisasecrettoken",
		RefreshToken: "thisisarefreshtoken",
		ExpiredAt:    time.Now().Add(time.Duration(a.config.TokenRepository.TokenExpiration) * time.Hour),
	}, nil
}

func (a *tokenRepository) Update(ctx context.Context, refreshToken string) (entity.Token, error) {
	return entity.Token{}, nil
}

func (a *tokenRepository) Delete(ctx context.Context, token string) error {
	return nil
}

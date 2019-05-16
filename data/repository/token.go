package repository

import (
	"database/sql"
	"time"

	"github.com/go-redis/redis"
	"github.com/syafdia/go-clean-arch/data/entity"
)

type TokenRepository interface {
	Generate(userID string, username string) (entity.Token, error)
	Update(token string) (entity.Token, error)
	Delete(token string) error
}

type tokenRepository struct {
	db          *sql.DB
	redisClient *redis.Client
}

func NewtokenRepository(db *sql.DB, redisClient *redis.Client) TokenRepository {
	return &tokenRepository{db, redisClient}
}

func (a *tokenRepository) Generate(userID string, username string) (entity.Token, error) {
	return entity.Token{
		UserID:       userID,
		Value:        "thisisasecrettoken",
		RefreshToken: "thisisarefreshtoken",
		ExpiredAt:    time.Now().Add(12 * time.Hour),
	}, nil
}

func (a *tokenRepository) Update(refreshToken string) (entity.Token, error) {
	return entity.Token{}, nil
}

func (a *tokenRepository) Delete(token string) error {
	return nil
}

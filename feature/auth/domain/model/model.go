package model

import (
	"time"

	"github.com/syafdia/go-clean-arch/data/entity"
)

type InputSignIn struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type ResponseSignIn struct {
	UserID         string          `json:"user_id"`
	Username       string          `json:"username"`
	Role           entity.UserRole `json:"role"`
	Token          string          `json:"token"`
	RefreshToken   string          `json:"refresh_token"`
	TokenExpiredAt time.Time       `json:"token_expired_at"`
}

type InputSignUp struct {
}

var (
	EmptyResponseSignIn ResponseSignIn
)

package entity

import (
	"time"
)

type SignInRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type SignInResponse struct {
	UserID         string    `json:"user_id"`
	Username       string    `json:"username"`
	Role           UserRole  `json:"role"`
	Token          string    `json:"token"`
	RefreshToken   string    `json:"refresh_token"`
	TokenExpiredAt time.Time `json:"token_expired_at"`
}

type SignUpRequest struct {
	Username             string `form:"username" json:"username" binding:"required"`
	Password             string `form:"password" json:"password" binding:"required"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" binding:"required"`
}

type SignUpResponse struct {
	UserID         string    `json:"user_id"`
	Username       string    `json:"username"`
	Role           UserRole  `json:"role"`
	Token          string    `json:"token"`
	RefreshToken   string    `json:"refresh_token"`
	TokenExpiredAt time.Time `json:"token_expired_at"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	Token          string    `json:"token"`
	RefreshToken   string    `json:"refresh_token"`
	TokenExpiredAt time.Time `json:"token_expired_at"`
}

type ResetPasswordResponse struct {
	ResetPasswordURL string    `json:"url"`
	ExpiredAt        time.Time `json:"expired_at"`
}

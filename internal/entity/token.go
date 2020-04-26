package entity

import (
	"errors"
	"time"
)

var (
	ErrFailedGeneratingToken = errors.New("failed generating user token")
)

type Token struct {
	UserID       string
	Value        string
	RefreshToken string
	ExpiredAt    time.Time
}

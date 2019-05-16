package entity

import (
	"time"
)

type Token struct {
	UserID       string
	Value        string
	RefreshToken string
	ExpiredAt    time.Time
}

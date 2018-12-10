package entity

import (
	"time"
)

type Token struct {
	Value     string
	ExpiredAt time.Time
}

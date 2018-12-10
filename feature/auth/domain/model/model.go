package model

import (
	"github.com/syafdia/go-clean-arch/data/entity"
)

type InputSignIn struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type ResponseSignIn struct {
	entity.User
}

type InputSignUp struct {
}

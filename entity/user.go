package entity

import "errors"

type UserRole string

const (
	UserRoleAdministrator UserRole = "administrator"
	UserRoleDefault       UserRole = "default"
)

var (
	ErrUserNotFound = errors.New("user is not found")
)

type User struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Role     UserRole `json:"role"`
}

package entity

type UserRole string

const (
	UserRoleAdministrator UserRole = "administrator"
	UserRoleDefault       UserRole = "default"
)

type User struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Role     UserRole `json:"role"`
}

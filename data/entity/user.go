package entity

const (
	UserRoleAdministrator = "administrator"
	UserRoleDefault       = "default"
)

type UserRole string

type User struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Role     UserRole `json:"role"`
}

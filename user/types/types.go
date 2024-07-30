package types

import "time"

type UserStore interface {
	ListUsers() ([]User, error)
	CreateUser(User) error
	GetUserById(int) (*User, error)
	GetUsersByEmail(string) ([]User, error)
	GetUsersByName(string) ([]User, error)
	UpdateUser(int, User) error
	DeleteUser(int) error
}

type UserRole string

const (
	Admin  UserRole = "admin"
	Client UserRole = "client"
)

type User struct {
	ID           int       `json:"id"`
	FullName     string    `json:"full_name"`
	Address      string    `json:"address"`
	Email        string    `json:"email"`
	RegisterDate time.Time `json:"register_date"`
	UserRole     UserRole  `json:"user_role"`
}

type CreateUserPayload struct {
	FullName string   `json:"full_name" validate:"required"`
	Address  string   `json:"address" validate:"required"`
	Email    string   `json:"email" validate:"required"`
	UserRole UserRole `json:"user_role" validate:"required"`
}

type UpdateUserPayload struct {
	FullName string   `json:"full_name" validate:"omitempty"`
	Address  string   `json:"address" validate:"omitempty"`
	UserRole UserRole `json:"user_role" validate:"omitempty"`
}

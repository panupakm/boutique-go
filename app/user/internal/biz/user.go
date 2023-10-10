package biz

import (
	"context"
)

type User struct {
	ID           string
	Name         string
	Email        string
	PasswordHash string
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Create(context.Context, *User) (*User, error)
	FindByID(context.Context, string) (*User, error)
	Update(context.Context, *User) (*User, error)
}

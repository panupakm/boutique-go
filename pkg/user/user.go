package user

import "github.com/google/uuid"

type User struct {
	Id           uuid.UUID
	Name         string
	Email        string
	PasswordHash string
}

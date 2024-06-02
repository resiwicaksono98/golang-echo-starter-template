package domain

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type UserRepository interface {
	FindAll() ([]User, error)
}

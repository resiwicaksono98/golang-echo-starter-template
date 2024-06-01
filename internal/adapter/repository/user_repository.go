package repository

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID
	Name string
}

type UserRepository interface {
	FindAll() ([]User, error)
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindAll() ([]User, error) {
	// Implementasi query untuk mendapatkan semua user
	return []User{{ID: uuid.New(), Name: "John Doe"}}, nil
}

package repository

import (
	"echo-starter-template/internal/domain"

	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID
	Name string
}

type userRepository struct {
}

func NewUserRepository() domain.UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindAll() ([]domain.User, error) {
	return []domain.User{{ID: uuid.New(), Name: "John Doe"}}, nil
}

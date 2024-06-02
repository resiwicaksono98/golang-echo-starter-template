package service

import (
	"echo-starter-template/internal/domain"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) FindAll() ([]domain.User, error) {
	return s.repo.FindAll()
}

package service

import "echo-starter-template/internal/adapter/repository"

type UserService interface {
	FindAll() ([]repository.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) FindAll() ([]repository.User, error) {
	return s.repo.FindAll()
}

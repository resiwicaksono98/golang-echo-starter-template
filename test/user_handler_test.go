package test

import (
	"echo-starter-template/internal/domain"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) FindAll() ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockService := new(MockUserService)
	expectedUsers := []domain.User{{ID: uuid.New(), Name: "John Doe"}}
	mockService.On("FindAll").Return(expectedUsers, nil)
	// Setup handler
	users, err := mockService.FindAll()

	assert.Nil(t, err)
	assert.Equal(t, expectedUsers, users)

	mockService.AssertExpectations(t)
}

package service

import (
	"backend/internal/dto"
	"backend/internal/model"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByEmail(email string) (*model.User, error) {
	args := m.Called(email)
	return nil, args.Error(1)
}

func TestRegister_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockRepo.On("Create", mock.Anything).Return(nil)

	userService := NewUserService(mockRepo)

	req := dto.RegisterRequestDTO{
		Email:    "test@test.com",
		Username: "testuser",
		Password: "password123",
	}

	err := userService.Register(req)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestRegister_DatabaseFailure(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockRepo.On("Create", mock.Anything).Return(errors.New("database connection lost"))

	userService := NewUserService(mockRepo)

	req := dto.RegisterRequestDTO{
		Email:    "test@test.com",
		Username: "testuser",
		Password: "password123",
	}

	err := userService.Register(req)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database connection lost")
	mockRepo.AssertExpectations(t)
}

func (m *MockUserRepository) DeleteAccount(userID string) error {
	args := m.Called(userID)
	return args.Error(0)
}

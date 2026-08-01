package service

import (
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/repository"
	"fmt"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(req dto.RegisterRequestDTO) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{userRepo: repo}
}

func (s *userService) Register(req dto.RegisterRequestDTO) error {
	//* BUssiness Logic
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	//! generate UUID
	userID := uuid.New().String()

	//! Mapping
	user := model.User{
		ID:       userID,
		Email:    req.Email,
		Username: req.Username,
		Password: string(hashedPassword),
	}

	if err := s.userRepo.Create(&user); err != nil {
		return fmt.Errorf("failed to save user record: %w", err)
	}

	fmt.Printf(`Processing registration in service layer for user: %s (%s)\n`, user.Username, user.Email)
	return nil
}

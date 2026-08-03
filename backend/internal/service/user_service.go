package service

import (
	"backend/internal/dto"
	"backend/internal/model"
	"backend/internal/repository"
	"backend/internal/utils"
	"fmt"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(req dto.RegisterRequestDTO) error
	Login(req dto.LoginRequestDTO) (string, error)
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

	fmt.Printf(`Successfully registered user entity: %s (%s)\n`, user.Username, user.Email)
	return nil
}

func (s *userService) Login(req dto.LoginRequestDTO) (string, error) {
	//! Locating User via Physical repository sequence
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return "", fmt.Errorf("User Not Found: %w", err)
	}

	//! Password Validation (plain-text against stored cryptographic hash)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", fmt.Errorf("Invalid Credentials: %w", err)
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to generate auth token: %w", err)
	}

	return token, nil
}

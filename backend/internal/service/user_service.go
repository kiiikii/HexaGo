package service

import (
	"backend/internal/dto"
	"backend/internal/repository"
	"fmt"
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
	s.userRepo.Create(req.Email, req.Username, req.Password)
	fmt.Printf(`Processing registration in service layer for user: %s (%s)\n`, req.Username, req.Email)
	return nil
}

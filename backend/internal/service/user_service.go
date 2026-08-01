package service

import (
	"backend/internal/dto"
	"fmt"
)

type UserService interface {
	Register(req dto.RegisterRequestDTO) error
}

type userService struct{}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) Register(req dto.RegisterRequestDTO) error {
	//* BUssiness Logic
	fmt.Printf(`Processing registration in service layer for user: %s (%s)\n`, req.Username, req.Email)
	return nil
}

package repository

import "fmt"

type UserRepository interface {
	Create(email string, username string, password string) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(email string, username string, password string) error {
	fmt.Println("Saved to Database")
	return nil
}

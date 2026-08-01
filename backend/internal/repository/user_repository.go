package repository

import (
	"backend/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) error {
	if result := r.db.Create(user); result.Error != nil {
		return result.Error
	}
	return nil
}

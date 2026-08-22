package repository

import (
	"backend/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	DeleteAccount(userID string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Create(user *model.User) error {
	if result := r.db.Create(user); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userRepository) DeleteAccount(userID string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		/**
		 *! Delete All message belonging to this user
		 *! using tx not r.db
		 */
		if err := tx.Where("user_id = ?", userID).Delete(&model.Message{}).Error; err != nil {
			//! Returning an error automatically TRIGGER A ROLLBACK
			return err
		}

		//! Delete User profile
		if err := tx.Where("id = ?", userID).Delete(&model.User{}).Error; err != nil {
			return err
		}

		return nil
	})
}

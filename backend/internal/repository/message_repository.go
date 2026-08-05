package repository

import (
	"backend/internal/model"

	"gorm.io/gorm"
)

type MessageRepository interface {
	SaveMessage(msg *model.Message) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) SaveMessage(msg *model.Message) error {
	return r.db.Create(msg).Error
}

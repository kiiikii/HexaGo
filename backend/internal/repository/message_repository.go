package repository

import (
	"backend/internal/model"

	"gorm.io/gorm"
)

type MessageRepository interface {
	SaveMessage(msg *model.Message) error
	GetMessage(limit int) ([]model.Message, error)
	GetMessages(limit int, offset int) ([]model.Message, error)
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{
		db: db,
	}
}

func (r *messageRepository) SaveMessage(msg *model.Message) error {
	return r.db.Create(msg).Error
}

func (r *messageRepository) GetMessage(limit int) ([]model.Message, error) {
	var message []model.Message
	if err := r.db.Preload("User").Order("created_at asc").Limit(limit).Find(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

func (r *messageRepository) GetMessages(limit int, offset int) ([]model.Message, error) {
	var messages []model.Message

	//! Order by latest first, preload the user struct, apply paggination limit
	err := r.db.Preload("User").Order("created_at desc").Limit(limit).Offset(offset).Find(&messages).Error
	return messages, err
}

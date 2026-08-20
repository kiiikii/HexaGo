package service

import (
	"backend/internal/model"
	"backend/internal/repository"

	"github.com/google/uuid"
)

type MessageService interface {
	SaveMessage(userID string, content string, room string) (*model.Message, error)
	GetMessage(limit int) ([]model.Message, error)
	GetMessages(limit int, offset int, room string) ([]model.Message, error)
	DeleteMessage(messageID string, userID string) error
}

type messageService struct {
	messageRepo repository.MessageRepository
}

func NewMessageService(repo repository.MessageRepository) MessageService {
	return &messageService{messageRepo: repo}
}

func (s *messageService) SaveMessage(userID string, content string, room string) (*model.Message, error) {
	messageID := uuid.New().String()

	msg := model.Message{
		ID:      messageID,
		UserID:  userID,
		Content: content,
		Room:    room,
	}

	if err := s.messageRepo.SaveMessage(&msg); err != nil {
		return nil, err
	}

	return &msg, nil
}

func (s *messageService) GetMessage(limit int) ([]model.Message, error) {
	return s.messageRepo.GetMessage(limit)
}

func (s *messageService) GetMessages(limit int, offset int, room string) ([]model.Message, error) {
	return s.messageRepo.GetMessages(limit, offset, room)
}

func (s *messageService) DeleteMessage(messageID string, userID string) error {
	return s.messageRepo.DeleteMessage(messageID, userID)
}

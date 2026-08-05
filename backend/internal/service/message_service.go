package service

import (
	"backend/internal/model"
	"backend/internal/repository"

	"github.com/google/uuid"
)

type MessageService interface {
	SaveMessage(userID string, content string) error
}

type messageService struct {
	messageRepo repository.MessageRepository
}

func NewMessageService(repo repository.MessageRepository) MessageService {
	return &messageService{messageRepo: repo}
}

func (s *messageService) SaveMessage(userID string, content string) error {
	messageID := uuid.New().String()

	msg := model.Message{
		ID:      messageID,
		UserID:  userID,
		Content: content,
	}

	if err := s.messageRepo.SaveMessage(&msg); err != nil {
		return err
	}

	return nil
}

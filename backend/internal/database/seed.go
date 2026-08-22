package database

import (
	"backend/internal/model"
	"log/slog"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	var count int64
	db.Model(&model.User{}).Count(&count)
	if count > 0 {
		slog.Info(" Database already seeded, skipping...")
		return
	}

	slog.Info("Seeding database with initial data...")

	//! Generate secure password for dummy
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)

	//! Creating User
	aliceID := uuid.New().String()
	alice := model.User{
		ID:       aliceID,
		Username: "Alice",
		Email:    "alice@hexago.com",
		Password: string(hashedPassword),
	}

	//! Save alice to DB
	if err := db.Create(&alice).Error; err != nil {
		slog.Error("Failed to seed user", slog.String("error", err.Error()))
		return
	}

	//! Create a welcome message
	welcomeMsg := model.Message{
		ID:      uuid.New().String(),
		UserID:  aliceID,
		Content: "Welcome to Hexago! this is an automatically generate message",
		Room:    "general",
	}

	db.Create(&welcomeMsg)
	slog.Info("Database seeding completed successfully")
}

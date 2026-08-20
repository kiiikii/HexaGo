package model

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	UserID    string         `gorm:"not null" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
	Content   string         `gorm:"not null" json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

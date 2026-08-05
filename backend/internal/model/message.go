package model

import (
	"time"
)

type Message struct {
	ID        string `gorm:"primaryKey"`
	UserID    string `gorm:"foreignKey:UserID;references:ID"`
	Content   string `gorm:"not null"`
	CreatedAt time.Time
}

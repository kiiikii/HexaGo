package model

import (
	"time"
)

type Message struct {
	ID        string `gorm:"primaryKey"`
	UserID    User   `gorm:"foreignKey:UserID;references:ID"`
	Content   string `gorm:"not null"`
	CreatedAr time.Time
}

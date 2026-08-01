package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	Email     string `gorm:"unique;not null"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

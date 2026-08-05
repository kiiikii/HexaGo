package database

import (
	"backend/internal/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	//! Logic Goes Here
	dsn := "host=localhost user=postgres password=secret dbname=hexago_chat port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Connect to database: ", err)
	}

	db.AutoMigrate(&model.User{}, &model.Message{})

	return db
}

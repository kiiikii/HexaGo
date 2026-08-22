package database

import (
	"backend/internal/model"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	//! Logic Goes Here
	dsn := "host=db user=postgres password=secret dbname=hexago_chat port=5432 sslmode=disable"

	var db *gorm.DB
	var err error

	//! Retry Loop
	for i := 1; i <= 5; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Fatal("Failed to Connect to database: ", err)
			break
		}

		log.Printf("Database not ready (attempt %d/5). Retryingin 2 sec...\n", i)
		time.Sleep(2 * time.Second)

	}

	if err != nil {
		log.Fatal(" Couldn't connect to database after 5 attemps: ", err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Message{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	Seed(db)

	return db
}

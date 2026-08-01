package main

import (
	"backend/internal/database"
	"backend/internal/router"
	"log"
)

func main() {
	//! Database Initialize
	db := database.Connect()

	//! Log for Testing
	log.Printf("Successfuly connected to database instance: %v", db)

	//! Router
	r := router.SetupRouter(db)
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

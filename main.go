package main

import (
	"log"
	"taskmanager-api/database"
	"taskmanager-api/models"
	"taskmanager-api/routes"

	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.Connect()

	database.DB.AutoMigrate(&models.Task{})

	router := routes.SetupRouter()
	router.Run(":8080")
}

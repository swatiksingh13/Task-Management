package main

import (
	"Task_Manager/database"
	"Task_Manager/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	log.Println("Connecting to the database...")
	database.Connect()
	defer func() {
		log.Println("Closing the database connection...")
		database.Close()
	}()

	log.Println("Initializing Gin router...")
	router := gin.Default()

	log.Println("Setting up CORS middleware...")
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	log.Println("Setting up routes...")
	routes.SetupRoutes(router)

	log.Println("Starting server on port 8081 🚀")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

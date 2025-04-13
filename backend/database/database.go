package database

import (
	"Task_Manager/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var DB *gorm.DB

func Connect() {
	var err error
	log.Println("Connecting to SQLite database...")
	DB, err = gorm.Open("sqlite3", "./taskmanager.db")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	log.Println("Connected successfully !!")
	log.Println("Running migrations")
	DB.AutoMigrate(&models.Task{})
	log.Println("Migration complete")
}

func Close() {
	log.Println("Closing database connection.")
	DB.Close()
	log.Println("Connection closed.")
}

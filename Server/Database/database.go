package database

import (
	"log"
	"os"
	models "server/Models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	username := os.Getenv("databaseUsername")
	password := os.Getenv("databasePassword")
	databaseName := os.Getenv("databaseName")
	dsn := username + ":" + password + "@/" + databaseName
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	DB = db
	db.AutoMigrate(&models.Blog{}, &models.User{})
}

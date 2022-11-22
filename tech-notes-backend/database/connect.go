package database

import (
	"fmt"
	"log"
	"os"
	"tech-notes-backend/configs"
	"tech-notes-backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB gorm connector
var DB *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var err error
	user := configs.Config("DB_USER")
	password := configs.Config("DB_PASSWORD")
	host := configs.Config("DB_HOST")
	port := configs.Config("DB_PORT")
	dbName := configs.Config("DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName,
	)

	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal("Failed to connect to database: ", err.Error())
		os.Exit(1)
	}

	log.Println("Connected to database")

	DB.AutoMigrate(&models.Role{}, &models.User{}, &models.Note{})

	log.Println("Database Migrated")
}

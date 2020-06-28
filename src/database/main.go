package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"go-to-do/src/models"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Database string
	Password string
}

func InitializeDatabase() *gorm.DB {

	config := DatabaseConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "postgres",
		Password: "postgres",
		Database:   "postgres",
	}
	var urlString = fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=disable password=%s",
		config.Host,
		config.User,
		config.Database,
		config.Password,
	)
	
	database, error := gorm.Open("postgres", urlString)

	if error != nil {
		fmt.Println("Database connection failed:", error)
	}
	
	fmt.Println("Migrating models...")
	database.AutoMigrate(&models.Todo{})

	fmt.Println("Successfully connected to database")
	
	return database
}
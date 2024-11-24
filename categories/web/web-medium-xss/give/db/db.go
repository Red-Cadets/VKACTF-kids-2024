package db

import (
	"example/v3/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenDBConnection() (*gorm.DB, error) {

	if DB != nil {
		return DB, nil
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to DB: %v", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Appointments{})
	if err != nil {
		return nil, fmt.Errorf("error during migration: %v", err)
	}

	fmt.Println("Success connection to DB and migrations completed")
	DB = db
	return DB, nil

}

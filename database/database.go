package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

func Connect() {
	var err error
	host := os.Getenv("DB_HOST")
	userName := os.Getenv("DB_USERNAME")
	name := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	port := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Cairo",
		host, userName, password, name, port)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Database Successfuly Connected!")
	}
}

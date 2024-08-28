package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env: %v", err)
	}
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	name := os.Getenv("MYSQL_DATABASE")

	dsn := user + ":" + password + "@tcp(" + host + ":3306)/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

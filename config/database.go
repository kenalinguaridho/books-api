package config

import (
	"github.com/kenalinguaridho/books-api/entities"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnect() {
	dsn := "root:@tcp(127.0.0.1:3306)/library_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entities.Author{}, &entities.Book{}, &entities.User{})

	log.Println("Database Connected!")
	DB = db
}
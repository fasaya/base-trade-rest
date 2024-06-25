package database

import (
	"base-trade-rest/core/model"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {

	// Load env
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file while starting db")
	// }

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbport := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	config := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbport, dbname)
	db, err = gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Debug().AutoMigrate(model.User{}, model.Product{}, model.Variant{})
}

func GetDB() *gorm.DB {
	return db
}

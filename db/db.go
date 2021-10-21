package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	uri := "host=localhost port=5001 database=books sslmode=disable user=admin password=123456"

	database, err := gorm.Open(postgres.Open(uri), &gorm.Config{})

	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	db = database
	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func GetDatabase() *gorm.DB {
	return db
}

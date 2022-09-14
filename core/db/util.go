package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	database, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})

	if err != nil {
		panic("DB connection failed")
	}

	database.AutoMigrate(&Short{})

	DB = database
}

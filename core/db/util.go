package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dbPath string) {
	database, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		panic("DB connection failed")
	}

	database.AutoMigrate(&User{})
	database.AutoMigrate(&Short{})

	DB = database
}

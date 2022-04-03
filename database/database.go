package database

import (
	"sample/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("skelp.db"))
	if err != nil {
		panic("Error")
	}

	db.AutoMigrate(&models.User{})
	return db
}
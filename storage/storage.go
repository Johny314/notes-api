package storage

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/jeka314/notes-api/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// Миграция схемы
	err = DB.AutoMigrate(&models.Note{})
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}
}

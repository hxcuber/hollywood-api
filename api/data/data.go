package data

import (
	"github.com/hxcuber/hollywood-api/api/internal/model"
	"gorm.io/gorm"
	"log"
	"os"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(model.Actor{})
	db.AutoMigrate(model.Movies{})
}

func InsertMockData(db *gorm.DB) {
	cActors, err := os.ReadFile("./migrations/actors.sql")

	if err != nil {
		log.Fatal(err)
	}
	sqlActors := string(c)
}

package db

import (
	"github.com/cottonBlanket/country-counter/internal/models/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init(dsn string) {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	err = DB.AutoMigrate(
		&domain.User{},
		&domain.Country{},
		&domain.Trip{},
	)
	if err != nil {
		log.Fatalf("auto migration failed: %v", err)
	}
}

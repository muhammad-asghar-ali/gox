package database

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"authis/internal/config"
)

var (
	instance *gorm.DB
	once     sync.Once
)

func Connect(cfg *config.Config) *gorm.DB {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
		)

		var err error
		instance, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Could not connect to database!", err)
		}
		log.Println("Database connection established successfully")
	})

	return instance
}

func GetDB() *gorm.DB {
	if instance == nil {
		log.Fatal("Database connection has not been initialized. Call Connect() first.")
	}

	return instance
}

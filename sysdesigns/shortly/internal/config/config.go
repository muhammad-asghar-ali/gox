package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type (
	Database struct {
		URI          string
		DatabaseName string
		Username     string
		Password     string
	}

	Server struct {
		Port string
	}

	Config struct {
		Server   Server
		Database Database
	}
)

var (
	instance *Config
	once     sync.Once
)

func Load() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}

		database := Database{
			URI:          os.Getenv("DATABASE_URI"),
			DatabaseName: os.Getenv("DATABASE_NAME"),
			Username:     os.Getenv("DATABASE_USERNAME"),
			Password:     os.Getenv("DATABASE_PASSWORD"),
		}

		server := Server{
			Port: os.Getenv("PORT"),
		}

		instance = &Config{
			Database: database,
			Server:   server,
		}
	})

	return instance
}

func GetConfig() *Config {
	if instance == nil {
		Load()
	}

	return instance
}

func (c *Config) GetDatabaseName() string {
	return instance.Database.DatabaseName
}

func (c *Config) GetURI() string {
	return instance.Database.URI
}

func (c *Config) GetPort() string {
	return instance.Server.Port
}

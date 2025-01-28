package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

	Redis struct {
		RedisAddr     string
		RedisPassword string
		RedisDB       int
	}

	Server struct {
		Port string
	}

	Config struct {
		Server   Server
		Database Database
		Redis    Redis
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

		redis := Redis{
			RedisAddr:     os.Getenv("REDIS_ADDR"),
			RedisPassword: os.Getenv("REDIS_PASSWORD"),
			RedisDB:       getRedisDB(),
		}

		server := Server{
			Port: os.Getenv("PORT"),
		}

		instance = &Config{
			Database: database,
			Server:   server,
			Redis:    redis,
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

func (c *Config) GetConnectionURL() string {
	return fmt.Sprintf("mongodb://%s:%s@localhost/%s?authSource=admin",
		instance.Database.Username,
		instance.Database.Password,
		instance.Database.DatabaseName,
	)
}

func (c *Config) GetRedisAddr() string {
	return c.Redis.RedisAddr
}

func (c *Config) GetRedisPassword() string {
	return c.Redis.RedisPassword
}

func (c *Config) GetRedisDB() int {
	return c.Redis.RedisDB
}

func getRedisDB() int {
	db := os.Getenv("REDIS_DB")
	if db == "" {
		return 0
	}

	value, err := strconv.Atoi(db)
	if err != nil {
		log.Fatalf("Invalid REDIS_DB value: %v", err)
	}

	return value
}

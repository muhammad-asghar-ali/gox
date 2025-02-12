package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		DBHost        string
		DBUser        string
		DBPassword    string
		DBName        string
		DBPort        string
		RedisAddr     string
		RedisPassword string
		RedisDB       int
		JWTSecret     string
	}
)

var (
	instance *Config
	once     sync.Once
)

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{
		DBHost:        os.Getenv("DB_HOST"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		DBPort:        os.Getenv("DB_PORT"),
		RedisAddr:     os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       getRedisDB(),
		JWTSecret:     os.Getenv("JWT_SECRET"),
	}
}

func GetConfig() *Config {
	once.Do(func() {
		instance = LoadConfig()
	})

	return instance
}

func (c *Config) GetRedisAddr() string {
	return c.RedisAddr
}

func (c *Config) GetRedisPassword() string {
	return c.RedisPassword
}

func (c *Config) GetRedisDB() int {
	return c.RedisDB
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

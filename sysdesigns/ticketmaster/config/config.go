package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

type (
	Database struct {
		Host      string
		Port      int
		User      string
		Password  string
		Name      string
		EnableSSL bool
	}

	Server struct {
		Port string
		Jwt  string
	}

	Redis struct {
		RedisAddr     string
		RedisPassword string
		RedisDB       int
	}

	Config struct {
		Server   Server
		Database Database
		Redis    Redis
	}
)

var (
	c  *Config
	oc sync.Once
)

func Load() *Config {
	oc.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file")
		}

		db_port, _ := strconv.Atoi(os.Getenv("DATABASE_PORT"))
		db := Database{
			Host:      os.Getenv("DATABASE_HOST"),
			Port:      db_port,
			User:      os.Getenv("DATABASE_USER"),
			Password:  os.Getenv("DATABASE_PASSWORD"),
			Name:      os.Getenv("DATABASE_NAME"),
			EnableSSL: strings.ToLower(os.Getenv("DATABASE_SSL")) == "true",
		}

		redis := Redis{
			RedisAddr:     os.Getenv("REDIS_ADDR"),
			RedisPassword: os.Getenv("REDIS_PASSWORD"),
			RedisDB:       getRedisDB(),
		}

		server := Server{
			Port: os.Getenv("PORT"),
			Jwt:  os.Getenv("JWT_SECRET"),
		}

		c = &Config{
			Database: db,
			Server:   server,
			Redis:    redis,
		}
	})

	return c
}

func GetConfig() *Config {
	if c == nil {
		return Load()
	}

	return c
}

func (c *Config) GetDatabaseName() string {
	return c.Database.Name
}

func (c *Config) GetDatabaseHost() string {
	return c.Database.Host
}

func (c *Config) GetDatabasePort() int {
	return c.Database.Port
}

func (c *Config) GetDatabaseUser() string {
	return c.Database.User
}

func (c *Config) GetDatabasePassword() string {
	return c.Database.Password
}

func (c *Config) GetServerPort() string {
	return c.Server.Port
}

func (c *Config) ConnectionURI() string {
	ssl := "disable"
	if c.Database.EnableSSL {
		ssl = "require"
	}

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
		ssl,
	)
}

func (c *Config) GetJwtSecret() string {
	return c.Server.Jwt
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

func (c *Config) GetRedisAddr() string {
	return c.Redis.RedisAddr
}

func (c *Config) GetRedisPassword() string {
	return c.Redis.RedisPassword
}

func (c *Config) GetRedisDB() int {
	return c.Redis.RedisDB
}

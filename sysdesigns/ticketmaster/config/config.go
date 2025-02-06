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

	Config struct {
		Server   Server
		Database Database
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

		server := Server{
			Port: os.Getenv("PORT"),
			Jwt:  os.Getenv("JWT_SECRET"),
		}

		c = &Config{
			Database: db,
			Server:   server,
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

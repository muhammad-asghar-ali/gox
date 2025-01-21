package helpers

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var (
	cfg  *Config
	once sync.Once
)

type (
	Config struct {
		Host     string `json:"host" env:"HOST"`
		Port     int    `json:"port" env:"PORT"`
		User     string `json:"user" env:"USER"`
		Password string `json:"password" env:"PASSWORD"`
		Database string `json:"database" env:"DB_NAME"`
		SslMode  string `json:"ssl_mode" env:"SSL_MODE"`
	}
)

func LoadEnv() *Config {
	once.Do(func() {
		err := godotenv.Load(".env")
		HandleError(err)

		port, err := strconv.Atoi(os.Getenv("PORT"))
		HandleError(err)

		cfg = &Config{
			Host:     os.Getenv("HOST"),
			Port:     port,
			User:     os.Getenv("USER"),
			Password: os.Getenv("PASSWORD"),
			Database: os.Getenv("DB_NAME"),
			SslMode:  os.Getenv("SSL_MODE"),
		}
	})

	return cfg
}

func GetConfig() *Config {
	if cfg == nil {
		return LoadEnv()
	}

	return cfg
}

func ConnectionString() string {
	c := GetConfig()

	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Database, c.Password, c.SslMode)
}

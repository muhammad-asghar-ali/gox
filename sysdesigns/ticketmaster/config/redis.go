package config

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	rdb     *redis.Client
	oncerdb sync.Once
)

func ConnectRedis() {
	oncerdb.Do(func() {
		cfg := GetConfig()
		rdb = redis.NewClient(&redis.Options{
			Addr:     cfg.GetRedisAddr(),
			Password: cfg.GetRedisPassword(),
			DB:       cfg.GetRedisDB(),
		})

		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err := rdb.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalf("Failed to connect to Redis: %v", err)
		}
	})
}

func GetRedisClient() *redis.Client {
	if rdb == nil {
		ConnectRedis()
	}

	return rdb
}

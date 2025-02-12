package database

import (
	"authis/internal/config"
	"sync"

	"github.com/go-redis/redis"
)

var (
	rdb     *redis.Client
	oncerdb sync.Once
)

func ConnectRedis() {
	once.Do(func() {
		cfg := config.GetConfig()
		rdb = redis.NewClient(&redis.Options{
			Addr:     cfg.GetRedisAddr(),
			Password: cfg.GetRedisPassword(),
			DB:       cfg.GetRedisDB(),
		})
	})
}

func GetRedisClient() *redis.Client {
	if rdb == nil {
		ConnectRedis()
	}

	return rdb
}

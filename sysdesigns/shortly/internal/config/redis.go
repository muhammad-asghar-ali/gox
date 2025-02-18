package config

import (
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

var (
	rdb     *redis.Client
	oncerdb sync.Once
	KEY     = "global_counter"
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

		_, err := rdb.Ping().Result()
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

func IncrementCounter(ctx context.Context) (int64, error) {
	value, err := GetRedisClient().Incr(KEY).Result()
	if err != nil {
		return 0, err
	}

	return value, nil
}

func GetCounter(ctx context.Context) (int64, error) {
	value, err := GetRedisClient().Get(KEY).Int64()
	if err != nil {
		if err == redis.Nil {
			return 0, nil
		}

		return 0, err
	}

	return value, nil
}

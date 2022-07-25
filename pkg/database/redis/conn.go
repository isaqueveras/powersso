package redis

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/isaqueveras/power-sso/config"
)

// NewRedisClient returns new redis client
func NewRedisClient(cfg *config.Config) *redis.Client {
	if cfg.Redis.RedisAddr == "" {
		cfg.Redis.RedisAddr = ":6379"
	}

	return redis.NewClient(&redis.Options{
		Addr:         cfg.Redis.RedisAddr,
		MinIdleConns: cfg.Redis.MinIdleConns,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  time.Duration(cfg.Redis.PoolTimeout) * time.Second,
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,
	})
}

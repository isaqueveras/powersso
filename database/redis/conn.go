// Copyright (c) 2022 Isaque Veras
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package redis

// NewRedisClient returns new redis client
// func NewRedisClient(cfg *config.Config) *redis.Client {
// 	if cfg.Redis.RedisAddr == "" {
// 		cfg.Redis.RedisAddr = ":6379"
// 	}

// 	return redis.NewClient(&redis.Options{
// 		Addr:         cfg.Redis.RedisAddr,
// 		MinIdleConns: cfg.Redis.MinIdleConns,
// 		PoolSize:     cfg.Redis.PoolSize,
// 		PoolTimeout:  time.Duration(cfg.Redis.PoolTimeout) * time.Second,
// 		Password:     cfg.Redis.Password,
// 		DB:           cfg.Redis.DB,
// 	})
// }

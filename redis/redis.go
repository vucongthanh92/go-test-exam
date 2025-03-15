package redis

import (
	"github.com/redis/go-redis/v9"
	redisUtils "github.com/vucongthanh92/go-base-utils/redis"
	"github.com/vucongthanh92/go-test-exam/config"
)

type Client redis.UniversalClient

func Open(cfg *config.RedisConfig) Client {
	return redisUtils.NewUniversalRedisClient(redisUtils.Config{
		Addrs:    cfg.Addrs,
		Password: cfg.Password,
		Username: cfg.Username,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
}

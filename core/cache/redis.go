package cache

import (
	"IMChat/core/config"

	"github.com/redis/go-redis/v9"
)

func InitRedis(conf *config.Redis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     conf.Address,
		Password: conf.Password,
		DB:       conf.DB,
	})
}

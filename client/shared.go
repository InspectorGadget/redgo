package client

import (
	redis "github.com/redis/go-redis/v9"
)

func GetRedisClient() *redis.Client {
	return redis.NewClient(
		&redis.Options{
			Addr: ":6379",
		},
	)
}

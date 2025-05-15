package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func GetValue(client *redis.Client, key string) (string, error) {
	val, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return "", fmt.Errorf("key '%v' not found", key)
	}

	return val, nil
}

func SetValue(client *redis.Client, key string, value any) error {
	err := client.Set(context.Background(), key, value, 1*time.Hour).Err()
	if err != nil {
		return fmt.Errorf("an error has occurred: %v", err.Error())
	}

	return nil
}

package cache

import (
	"context"
	"strconv"

	"github.com/leandro-d-santos/no-code-api/config"
	internalLogger "github.com/leandro-d-santos/no-code-api/internal/logger"
	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	ctx    context.Context        = context.Background()
	logger *internalLogger.Logger = internalLogger.NewLogger("Redis")
)

func InitializeRedis() {
	db, _ := strconv.Atoi(config.Env.RedisDb)
	client = redis.NewClient(&redis.Options{
		Addr:     config.Env.RedisHost,
		Password: config.Env.RedisPassword,
		DB:       db,
	})
}

func Get(key string) (string, error) {
	value, err := client.Get(ctx, key).Result()
	if err != nil {
		logger.ErrorF("Error to get from cache. Key: %v", key)
		return "", err
	}
	return value, nil
}

func Set(key string, value interface{}) error {
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		logger.ErrorF("Error to set on cache. Key: %v / Value: %v", key, value)
	}
	return err
}

func Delete(key ...string) error {
	err := client.Del(ctx, key...).Err()
	if err != nil {
		logger.ErrorF("Error to delete from cache. Key: %v", key)
	}
	return err
}

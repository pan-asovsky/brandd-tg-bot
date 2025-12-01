package cache

import (
	"context"
	"fmt"

	"github.com/pan-asovsky/brandd-tg-bot/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewRedis(cfg *config.Config) (*redis.Client, error) {
	opt := &redis.Options{
		Addr:     cfg.RedisURL,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	}

	client := redis.NewClient(opt)
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("[redis_client] failed to connect to redis: %w", err)
	}

	return client, nil
}

package cache

import (
	"context"
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/config"
	"github.com/redis/go-redis/v9"
)

type Client struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedis(cfg *config.Config) (*Client, error) {
	ctx := context.Background()

	opt := &redis.Options{
		Addr:     cfg.RedisURL,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	}

	client := redis.NewClient(opt)

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &Client{client: client, ctx: ctx}, nil
}

func (r *Client) Set(key string, value []byte, ttl time.Duration) error {
	return r.client.Set(r.ctx, key, value, ttl).Err()
}

func (r *Client) Get(key string) ([]byte, error) {
	s, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return []byte(s), nil
}

func (r *Client) Del(key string) error {
	return r.client.Del(r.ctx, key).Err()
}

func (r *Client) Close() error {
	return r.client.Close()
}

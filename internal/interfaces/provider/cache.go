package provider

import (
	"time"

	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/cache"
	"github.com/redis/go-redis/v9"
)

type CacheProvider interface {
	SlotLock() i.SlotLockCache
	ServiceType() i.ServiceTypeCache
	RedisClient() *redis.Client
	TTL() time.Duration
}

package provider

import (
	"time"

	icache "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/cache"
	"github.com/redis/go-redis/v9"
)

type CacheProvider interface {
	SlotLock() icache.SlotLockCache
	ServiceType() icache.ServiceTypeCache
	RedisClient() *redis.Client
	TTL() time.Duration
}

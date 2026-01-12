package cache

import (
	"time"

	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/cache"
	p "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	"github.com/redis/go-redis/v9"
)

type cacheProvider struct {
	rc       *redis.Client
	cacheTTL time.Duration
}

func NewCacheProvider(rc *redis.Client, cacheTTL time.Duration) p.CacheProvider {
	return &cacheProvider{rc: rc, cacheTTL: cacheTTL}
}

func (p *cacheProvider) SlotLock() i.SlotLockCache {
	return NewSlotLockCache(p.rc, p.cacheTTL)
}

func (p *cacheProvider) ServiceType() i.ServiceTypeCache {
	return NewServiceTypeCacheService(p.rc, p.cacheTTL)
}

func (p *cacheProvider) RedisClient() *redis.Client {
	return p.rc
}

func (p *cacheProvider) TTL() time.Duration {
	return p.cacheTTL
}

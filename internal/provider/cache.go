package provider

import (
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/cache"
	icache "github.com/pan-asovsky/brandd-tg-bot/internal/interface/cache"
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interface/provider"
	"github.com/redis/go-redis/v9"
)

type cacheProvider struct {
	rc       *redis.Client
	cacheTTL time.Duration
}

func NewCacheProvider(rc *redis.Client, cacheTTL time.Duration) iprovider.CacheProvider {
	return &cacheProvider{rc: rc, cacheTTL: cacheTTL}
}

func (p *cacheProvider) SlotLock() icache.SlotLockCache {
	return cache.NewSlotLockCache(p.rc, p.cacheTTL)
}

func (p *cacheProvider) ServiceType() icache.ServiceTypeCache {
	return cache.NewServiceTypeCacheService(p.rc, p.cacheTTL)
}

func (p *cacheProvider) RedisClient() *redis.Client {
	return p.rc
}

func (p *cacheProvider) TTL() time.Duration {
	return p.cacheTTL
}

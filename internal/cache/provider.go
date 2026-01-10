package cache

import (
	"time"

	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/cache"
	"github.com/redis/go-redis/v9"
)

type Provider struct {
	rc       *redis.Client
	cacheTTL time.Duration
}

func NewProvider(rc *redis.Client, cacheTTL time.Duration) *Provider {
	return &Provider{rc: rc, cacheTTL: cacheTTL}
}

func (p *Provider) SlotLock() i.SlotLockCache {
	return NewSlotLockCache(p.rc, p.cacheTTL)
}

func (p *Provider) ServiceType() i.ServiceTypeCache {
	return NewServiceTypeCacheService(p.rc, p.cacheTTL)
}

func (p *Provider) RedisClient() *redis.Client {
	return p.rc
}

func (p *Provider) TTL() time.Duration {
	return p.cacheTTL
}

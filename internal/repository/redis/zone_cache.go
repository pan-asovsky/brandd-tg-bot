package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/cache"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/redis/go-redis/v9"
)

type ZoneCache struct {
	cache *cache.Client
	ttl   time.Duration
}

func NewZoneCache(r *cache.Client, ttl time.Duration) *ZoneCache {
	return &ZoneCache{
		cache: r,
		ttl:   ttl,
	}
}

func (c *ZoneCache) CacheZones(key string, zones model.Zone) error {
	data, err := json.Marshal(zones)
	if err != nil {
		return fmt.Errorf("failed to marshal zones: %w", err)
	}

	if err := c.cache.Set(key, data, c.ttl); err != nil {
		return fmt.Errorf("failed to set key in redis: %w", err)
	}

	log.Printf("[cache_zones] ok. key: %s ttl: %s", key, c.ttl)
	return nil
}

func (c *ZoneCache) GetZones(key string) (model.Zone, error) {
	data, err := c.cache.Get(key)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get key from redis: %w", err)
	}

	var zones model.Zone
	if err := json.Unmarshal(data, &zones); err != nil {
		return nil, fmt.Errorf("failed to unmarshal zones: %w", err)
	}

	log.Printf("[get_zones] ok. key: %s ttl: %s", key, c.ttl)
	return zones, nil
}

func FormatKey(key string, data string) string {
	return fmt.Sprintf("%s:%s", key, data)
}

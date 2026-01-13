package cache

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	icache "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/cache"
	"github.com/pan-asovsky/brandd-tg-bot/internal/rules"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
	"github.com/redis/go-redis/v9"
)

type serviceTypeCache struct {
	cache *redis.Client
	ttl   time.Duration
	ctx   context.Context
}

func NewServiceTypeCacheService(r *redis.Client, ttl time.Duration) icache.ServiceTypeCache {
	return &serviceTypeCache{
		cache: r,
		ttl:   ttl,
		ctx:   context.Background(),
	}
}

func (s *serviceTypeCache) Toggle(chatID int64, clickedService string) (map[string]bool, error) {
	key := s.formatKey(chatID)

	currentMap, err := s.cache.HGetAll(s.ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, utils.WrapError(err)
	}

	selectedMap := make(map[string]bool)
	for service, value := range currentMap {
		selectedMap[service] = value == "true"
	}

	selectedMap[clickedService] = !selectedMap[clickedService]

	serviceRules := &rules.ServiceRules{}
	finalMap := serviceRules.Apply(selectedMap, clickedService)

	redisData := make(map[string]interface{})
	for service, selected := range finalMap {
		if selected {
			redisData[service] = "true"
		} else {
			redisData[service] = "false"
		}
	}

	pipeline := s.cache.Pipeline()
	pipeline.Del(s.ctx, key)

	if len(redisData) > 0 {
		pipeline.HSet(s.ctx, key, redisData)
	}

	pipeline.Expire(s.ctx, key, s.ttl)

	if _, err := pipeline.Exec(s.ctx); err != nil {
		return nil, utils.WrapError(err)
	}

	return finalMap, nil
}

func (s *serviceTypeCache) Clean(chatID int64) error {
	key := s.formatKey(chatID)

	exists, err := s.cache.Exists(s.ctx, key).Result()
	if err != nil {
		return utils.WrapError(err)
	}

	if exists > 0 {
		if err := s.cache.Del(s.ctx, key).Err(); err != nil {
			return utils.WrapError(err)
		}
		log.Printf("[service_cache] deleted for: %d (existed)", chatID)
	} else {
		log.Printf("[service_cache] nothing to delete for: %d", chatID)
	}

	return nil
}
func (s *serviceTypeCache) formatKey(chatID int64) string {
	return fmt.Sprintf("selected_services:%d", chatID)
}

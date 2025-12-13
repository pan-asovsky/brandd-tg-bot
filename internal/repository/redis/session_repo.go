package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/rules"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
	"github.com/redis/go-redis/v9"
)

type SessionRepo struct {
	cache *redis.Client
	ttl   time.Duration
	ctx   context.Context
}

func NewSessionRepo(r *redis.Client, ttl time.Duration) *SessionRepo {
	return &SessionRepo{
		cache: r,
		ttl:   ttl,
		ctx:   context.Background(),
	}
}

func (s *SessionRepo) Toggle(chatID int64, clickedService string) (map[string]bool, error) {
	key := s.selectedServicesKey(chatID)

	currentMap, err := s.cache.HGetAll(s.ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return nil, utils.WrapError(err)
	}

	selectedMap := make(map[string]bool)
	for service, value := range currentMap {
		selectedMap[service] = value == "true"
	}

	log.Printf("[toggle] current %#v, selected %#v", currentMap, selectedMap)

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

	log.Printf("[toggle] final %#v", finalMap)
	return finalMap, nil
}

func (s *SessionRepo) selectedServicesKey(chatID int64) string {
	return fmt.Sprintf("selected_services:%d", chatID)
}

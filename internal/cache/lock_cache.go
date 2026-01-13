package cache

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	icache "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/cache"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
	"github.com/redis/go-redis/v9"
)

type slotLockCache struct {
	rc  *redis.Client
	ttl time.Duration
	ctx context.Context
}

func NewSlotLockCache(rc *redis.Client, ttl time.Duration) icache.SlotLockCache {
	return &slotLockCache{
		rc:  rc,
		ttl: ttl,
		ctx: context.Background(),
	}
}

func (lc *slotLockCache) Set(chatID int64, info model.SlotLockInfo) {
	val := fmt.Sprintf("%s|%s", info.Key, info.UUID)
	if _, err := lc.rc.Set(lc.ctx, lc.fmtKey(chatID), val, lc.ttl).Result(); err != nil {
		log.Printf("[lock_cache] error ignore?: %v", err)
	}
}

func (lc *slotLockCache) Get(chatID int64) (model.SlotLockInfo, bool, error) {
	val, err := lc.rc.Get(lc.ctx, lc.fmtKey(chatID)).Result()
	if errors.Is(err, redis.Nil) {
		return model.SlotLockInfo{}, false, nil
	}

	if err != nil {
		return model.SlotLockInfo{}, false, fmt.Errorf("[lock_cache] get error: %w", err)
	}

	key, uuid, found := strings.Cut(val, "|")
	if !found {
		return model.SlotLockInfo{}, false, fmt.Errorf("[lock_cache] invalid format: %s. error: %w", val, err)
	}

	return model.SlotLockInfo{Key: key, UUID: uuid}, true, nil
}

func (lc *slotLockCache) Del(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return lc.rc.Del(lc.ctx, lc.fmtKey(chatID)).Err()
	})
}

func (lc *slotLockCache) fmtKey(chatID int64) string {
	return fmt.Sprintf("user:%d:lock", chatID)
}

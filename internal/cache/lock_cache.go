package cache

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
	"github.com/redis/go-redis/v9"
)

type LockCache struct {
	rc  *redis.Client
	ttl time.Duration
	ctx context.Context
}

type SlotLockInfo struct {
	Key  string
	UUID string
}

func NewLockCache(rc *redis.Client, ttl time.Duration) *LockCache {
	return &LockCache{
		rc:  rc,
		ttl: ttl,
		ctx: context.Background(),
	}
}

func (lc *LockCache) Set(chatID int64, info SlotLockInfo) {
	val := fmt.Sprintf("%s|%s", info.Key, info.UUID)
	if _, err := lc.rc.Set(lc.ctx, lc.fmtKey(chatID), val, lc.ttl).Result(); err != nil {
		log.Printf("[lock_cache] error ignore?: %v", err)
	}
}

func (lc *LockCache) Get(chatID int64) (SlotLockInfo, bool, error) {
	val, err := lc.rc.Get(lc.ctx, lc.fmtKey(chatID)).Result()
	if errors.Is(err, redis.Nil) {
		return SlotLockInfo{}, false, nil
	}

	if err != nil {
		return SlotLockInfo{}, false, fmt.Errorf("[lock_cache] get error: %w", err)
	}

	key, uuid, found := strings.Cut(val, "|")
	if !found {
		return SlotLockInfo{}, false, fmt.Errorf("[lock_cache] invalid format: %s. error: %w", val, err)
	}

	return SlotLockInfo{Key: key, UUID: uuid}, true, nil
}

func (lc *LockCache) Del(chatID int64) error {
	return utils.WrapFunctionError(func() error {
		return lc.rc.Del(lc.ctx, lc.fmtKey(chatID)).Err()
	})
}

func (lc *LockCache) fmtKey(chatID int64) string {
	return fmt.Sprintf("user:%d:lock", chatID)
}

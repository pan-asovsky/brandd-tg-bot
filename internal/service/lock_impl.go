package service

import (
	"fmt"
	"log"

	"github.com/pan-asovsky/brandd-tg-bot/internal/cache"
	"github.com/pan-asovsky/brandd-tg-bot/internal/cache/locker"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
)

type lockService struct {
	locker *locker.SlotLocker
	cache  *cache.LockCache
}

func (ls *lockService) Toggle(info *types.UserSessionInfo) error {
	newKey := ls.locker.FormatKey(info.Date, info.Time)
	//log.Printf("[toggle] new key: %s", newKey)

	chatID := info.ChatID
	oldLock, exists, err := ls.cache.Get(chatID)
	if err != nil {
		return fmt.Errorf("[toggle] %w", err)
	}
	//log.Printf("[toggle] old lock exists: %v", exists)
	//log.Printf("[old_lock] key: %s, uuid: %s", oldLock.Key, oldLock.UUID)

	if exists && oldLock.Key == newKey {
		ls.cache.Set(chatID, oldLock)
		log.Printf("[extend_cache] new key: %s, old: %s", newKey, oldLock.Key)
		if err := ls.locker.RefreshTTL(oldLock.Key); err != nil {
			return fmt.Errorf("[refresh_ttl] error: %w", err)
		}
	}

	if exists {
		if err := ls.locker.Unlock(oldLock.Key, oldLock.UUID); err != nil {
			return fmt.Errorf("[toggle] %w", err)
		}
		ls.cache.Del(chatID)
	}

	uuid, ok, err := ls.locker.Lock(info.Date, info.Time)
	if err != nil || !ok {
		return fmt.Errorf("[toggle] %w", err)
	}

	lockInfo := cache.SlotLockInfo{Key: newKey, UUID: uuid}
	ls.cache.Set(chatID, lockInfo)

	return nil
}

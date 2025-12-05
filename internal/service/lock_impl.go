package service

import (
	"fmt"
	"log"

	"github.com/pan-asovsky/brandd-tg-bot/internal/cache"
	"github.com/pan-asovsky/brandd-tg-bot/internal/cache/locker"
)

type lockService struct {
	locker *locker.SlotLocker
	cache  *cache.LockCache
}

func (ls *lockService) Toggle(chatID int64, date, time string) error {
	newKey := ls.locker.FormatKey(date, time)
	//log.Printf("[toggle] new key: %s", newKey)

	oldLock, exists, err := ls.cache.Get(chatID)
	if err != nil {
		return fmt.Errorf("[old_lock] error: %w", err)
	}
	//log.Printf("[toggle] old lock exists: %v", exists)
	//log.Printf("[old_lock] key: %s, uuid: %s", oldLock.Key, oldLock.UUID)

	if exists && oldLock.Key == newKey {
		ls.cache.Set(chatID, oldLock)
		log.Printf("[extend_cache] new key: %s, old: %s", newKey, oldLock.Key)
		if err := ls.locker.RefreshTTL(oldLock.Key); err != nil {
			return fmt.Errorf("[refresh_ttl] error: %w", err)
		}
		return nil
	}

	if exists {
		if err := ls.locker.Unlock(oldLock.Key, oldLock.UUID); err != nil {
			return fmt.Errorf("[unlock] error: %w", err)
		}
		ls.cache.Del(chatID)
	}

	uuid, ok, err := ls.locker.Lock(date, time)
	if err != nil || !ok {
		return fmt.Errorf("[slot_locker] error: %w", err)
	}

	info := cache.SlotLockInfo{Key: newKey, UUID: uuid}
	//log.Printf("[toggle] new info: %v", info)
	ls.cache.Set(chatID, info)

	return nil
}

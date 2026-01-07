package service

import (
	"fmt"
	"log"

	"github.com/pan-asovsky/brandd-tg-bot/internal/cache"
	"github.com/pan-asovsky/brandd-tg-bot/internal/cache/locker"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type lockService struct {
	locker *locker.SlotLocker
	cache  *cache.LockCache
}

func (ls *lockService) Toggle(info *types.UserSessionInfo) error {
	newKey := ls.locker.FormatKey(info.Date, info.Time)

	chatID := info.ChatID
	oldLock, exists, err := ls.cache.Get(chatID)
	if err != nil {
		return fmt.Errorf("[time_lock] %w", err)
	}

	if exists && oldLock.Key == newKey {
		ls.cache.Set(chatID, oldLock)
		log.Printf("[time_lock] new key: %s, old: %s", newKey, oldLock.Key)
		if err := ls.locker.RefreshTTL(oldLock.Key); err != nil {
			return fmt.Errorf("[refresh_ttl] error: %w", err)
		}
	}

	if exists {
		if err := ls.locker.Unlock(oldLock.Key, oldLock.UUID); err != nil {
			return utils.WrapError(err)
		}
		if err := ls.cache.Del(chatID); err != nil {
			return utils.WrapError(err)
		}
	}

	uuid, ok, err := ls.locker.Lock(info.Date, info.Time)
	if err != nil || !ok {
		return fmt.Errorf("[time_lock] %w", err)
	}

	lockInfo := cache.SlotLockInfo{Key: newKey, UUID: uuid}
	ls.cache.Set(chatID, lockInfo)

	return nil
}

func (ls *lockService) Clean(chatID int64) error {
	lock, exists, err := ls.cache.Get(chatID)
	if err != nil {
		return utils.WrapError(err)
	}

	if exists {
		if err := ls.locker.Unlock(lock.Key, lock.UUID); err != nil {
			return fmt.Errorf("[time_lock] unlock error: %w", err)
		}
	}

	if err := ls.cache.Del(chatID); err != nil {
		return utils.WrapError(err)
	}

	log.Printf("[time_lock] deleted for: %d, exists: %v", chatID, exists)
	return nil
}

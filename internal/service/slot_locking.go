package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	"github.com/redis/go-redis/v9"
)

type slotLocking struct {
	rc        *redis.Client
	ttl       time.Duration
	lockSha   string
	unlockSha string
	ctx       context.Context
}

func NewSlotLocking(rc *redis.Client, ttl time.Duration) (i.SlotLocking, error) {
	lock, err := os.ReadFile("/app/script/lock.lua")
	if err != nil {
		return nil, fmt.Errorf("[slot_locker] failed to read lock.lua: %w", err)
	}
	unlock, err := os.ReadFile("/app/script/unlock.lua")
	if err != nil {
		return nil, fmt.Errorf("[slot_locker] failed to read unlock.lua: %w", err)
	}

	ctx := context.Background()
	lockSha, err := rc.ScriptLoad(ctx, string(lock)).Result()
	if err != nil {
		return nil, fmt.Errorf("[slot_locker] failed load lock.lua: %w", err)
	}
	unlockSha, err := rc.ScriptLoad(ctx, string(unlock)).Result()
	if err != nil {
		return nil, fmt.Errorf("[slot_locker] failed load unlock.lua: %w", err)
	}

	return &slotLocking{
		rc:        rc,
		ttl:       ttl,
		lockSha:   lockSha,
		unlockSha: unlockSha,
		ctx:       ctx,
	}, nil
}

func (sl *slotLocking) Lock(date, time string) (uid string, ok bool, err error) {
	key := sl.FormatKey(date, time)
	u := uuid.NewString()

	res, err := sl.rc.EvalSha(sl.ctx, sl.lockSha, []string{key}, u, sl.ttl.Milliseconds()).Result()
	if err != nil {
		return "", false, fmt.Errorf("[slot_locker] failed to lock: %w", err)
	}

	return u, res.(int64) == 1, nil
}

func (sl *slotLocking) Unlock(key, u string) error {
	res, err := sl.rc.EvalSha(sl.ctx, sl.unlockSha, []string{key}, u).Result()
	if err != nil {
		return err
	}
	if res.(int64) == 0 {
		return fmt.Errorf("lock not owned or already released")
	}
	return nil
}

func (sl *slotLocking) IsLocked(date, time string) (bool, error) {
	key := sl.FormatKey(date, time)
	val, err := sl.rc.Get(sl.ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return val != "", nil
}

func (sl *slotLocking) RefreshTTL(key string) error {
	_, err := sl.rc.PExpire(sl.ctx, key, sl.ttl).Result()
	return err
}

func (sl *slotLocking) AreLocked(keys ...string) (map[string]bool, error) {
	res, err := sl.rc.MGet(sl.ctx, keys...).Result()
	if err != nil {
		return nil, err
	}
	status := make(map[string]bool, len(keys))
	for x, key := range keys {
		status[key] = res[x] != nil
	}
	return status, nil
}

func (sl *slotLocking) FormatKey(date, time string) string {
	return fmt.Sprintf("s_lock:%s_%s", date, time)
}

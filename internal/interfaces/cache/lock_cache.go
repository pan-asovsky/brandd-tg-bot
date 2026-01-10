package cache

import "github.com/pan-asovsky/brandd-tg-bot/internal/model"

type SlotLockCache interface {
	Set(chatID int64, info model.SlotLockInfo)
	Get(chatID int64) (model.SlotLockInfo, bool, error)
	Del(chatID int64) error
}

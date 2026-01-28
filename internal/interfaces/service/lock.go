package service

import "github.com/pan-asovsky/brandd-tg-bot/internal/model"

type LockService interface {
	Toggle(info *model.UserSessionInfo) error
	Clean(chatID int64) error
}

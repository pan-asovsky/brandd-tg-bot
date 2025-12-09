package service

import "github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"

type LockService interface {
	Toggle(info *types.UserSessionInfo) error
}

package service

import "github.com/pan-asovsky/brandd-tg-bot/internal/entity"

type UserService interface {
	GetActiveAdmins() []entity.User
	GetRole(chatID int64) (bool, string)
}

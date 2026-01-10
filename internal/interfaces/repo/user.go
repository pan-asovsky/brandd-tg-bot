package interfaces

import "github.com/pan-asovsky/brandd-tg-bot/internal/entity"

type UserRepo interface {
	GetActiveAdmins() ([]entity.User, error)
}

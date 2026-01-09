package interfaces

import "github.com/pan-asovsky/brandd-tg-bot/internal/model"

type UserRepo interface {
	GetActiveAdmins() ([]model.User, error)
}

package service

import "github.com/pan-asovsky/brandd-tg-bot/internal/model"

type UserService interface {
	GetActiveAdmins() []model.User
}

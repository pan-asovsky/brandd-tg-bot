package service

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
)

type userService struct {
	repo i.UserRepo
}

func (us *userService) GetActiveAdmins() []entity.User {
	admins, err := us.repo.GetActiveAdmins()
	if err != nil {
		var empty []entity.User
		return empty
	}

	return admins
}

func (us *userService) GetRole(chatID int64) (bool, string) {
	return us.repo.GetRole(chatID)
}

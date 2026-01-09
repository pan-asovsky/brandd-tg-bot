package service

import (
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type userService struct {
	repo i.UserRepo
}

func (u *userService) GetActiveAdmins() []model.User {
	admins, err := u.repo.GetActiveAdmins()
	if err != nil {
		var empty []model.User
		return empty
	}

	return admins
}

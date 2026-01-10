package service

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
)

type userService struct {
	repo i.UserRepo
}

func (u *userService) GetActiveAdmins() []entity.User {
	admins, err := u.repo.GetActiveAdmins()
	if err != nil {
		var empty []entity.User
		return empty
	}

	return admins
}

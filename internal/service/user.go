package service

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	irepo "github.com/pan-asovsky/brandd-tg-bot/internal/interface/repo"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interface/service"
)

type userService struct {
	repo irepo.UserRepo
}

func NewUserService(userRepo irepo.UserRepo) isvc.UserService {
	return &userService{repo: userRepo}
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

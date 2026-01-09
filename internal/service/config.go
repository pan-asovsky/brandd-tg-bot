package service

import (
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type configService struct {
	configRepo i.ConfigRepo
}

func (c *configService) IsAutoConfirm() (bool, error) {
	return utils.WrapFunction(c.configRepo.IsAutoConfirm)
}

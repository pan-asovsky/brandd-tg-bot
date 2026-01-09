package service

import (
	pg "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type configService struct {
	configRepo pg.ConfigRepo
}

func (c *configService) IsAutoConfirm() (bool, error) {
	return utils.WrapFunction(c.configRepo.IsAutoConfirm)
}

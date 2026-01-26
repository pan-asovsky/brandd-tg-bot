package service

import (
	irepo "github.com/pan-asovsky/brandd-tg-bot/internal/interface/repo"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interface/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type configService struct {
	configRepo irepo.ConfigRepo
}

func NewConfigService(configRepo irepo.ConfigRepo) isvc.ConfigService {
	return &configService{configRepo}
}

func (c *configService) IsAutoConfirm() (bool, error) {
	return utils.WrapFunction(c.configRepo.IsAutoConfirm)
}

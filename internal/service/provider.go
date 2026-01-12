package service

import (
	p "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
)

type svcProvider struct {
	repoProvider  p.RepoProvider
	cacheProvider p.CacheProvider
}

func NewServiceProvider(repoProvider p.RepoProvider, cacheProvider p.CacheProvider) p.ServiceProvider {
	return &svcProvider{repoProvider: repoProvider, cacheProvider: cacheProvider}
}

func (sp *svcProvider) Slot() i.SlotService {
	return &slotService{sp.repoProvider.Slot(), sp.SlotLocking()}
}

func (sp *svcProvider) Keyboard() i.KeyboardService {
	return &keyboardService{callbackBuilding: sp.CallbackBuilding(), dateTime: sp.DateTime()}
}

func (sp *svcProvider) Lock() i.LockService {
	return &lockService{sp.SlotLocking(), sp.cacheProvider.SlotLock()}
}

func (sp *svcProvider) Booking() i.BookingService {
	return &bookingService{sp.repoProvider, sp.Slot(), sp.Price()}
}

func (sp *svcProvider) Price() i.PriceService {
	return &priceService{sp.repoProvider}
}

func (sp *svcProvider) Config() i.ConfigService {
	return &configService{sp.repoProvider.Config()}
}

func (sp *svcProvider) CallbackParsing() i.CallbackParsingService {
	return &callbackParsingService{}
}

func (sp *svcProvider) CallbackBuilding() i.CallbackBuildingService {
	return &callbackBuildingService{}
}

func (sp *svcProvider) DateTime() i.DateTimeService {
	return &dateTimeService{}
}

func (sp *svcProvider) User() i.UserService {
	return &userService{sp.repoProvider.User()}
}

func (sp *svcProvider) SlotLocking() i.SlotLocking {
	slotLock, err := NewSlotLocking(sp.cacheProvider.RedisClient(), sp.cacheProvider.TTL())
	if err != nil {
		panic(err)
	}
	return slotLock
}

func (sp *svcProvider) Phone() i.PhoneService {
	return NewPhoneNormalizingService()
}

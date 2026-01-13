package provider

import (
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service/callback"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service/keyboard"
)

type svcProvider struct {
	repoProvider  iprovider.RepoProvider
	cacheProvider iprovider.CacheProvider
}

func NewServiceProvider(repoProvider iprovider.RepoProvider, cacheProvider iprovider.CacheProvider) iprovider.ServiceProvider {
	return &svcProvider{repoProvider: repoProvider, cacheProvider: cacheProvider}
}

func (sp *svcProvider) Slot() isvc.SlotService {
	return service.NewSlotService(sp.repoProvider.Slot(), sp.SlotLocking())
}

func (sp *svcProvider) UserKeyboard() isvc.UserKeyboardService {
	return keyboard.NewUserKeyboardService(sp.CallbackBuilding(), sp.DateTime())
}

func (sp *svcProvider) AdminKeyboard() isvc.AdminKeyboardService {
	return keyboard.NewAdminKeyboardService(sp.CallbackBuilding(), sp.DateTime())
}

func (sp *svcProvider) Lock() isvc.LockService {
	return service.NewLockService(sp.SlotLocking(), sp.cacheProvider.SlotLock())
}

func (sp *svcProvider) Booking() isvc.BookingService {
	return service.NewBookingService(sp.repoProvider, sp.Slot(), sp.Price())
}

func (sp *svcProvider) Price() isvc.PriceService {
	return service.NewPriceService(sp.repoProvider.Price())
}

func (sp *svcProvider) Config() isvc.ConfigService {
	return service.NewConfigService(sp.repoProvider.Config())
}

func (sp *svcProvider) CallbackParsing() isvc.CallbackParsingService {
	return callback.NewCallbackParsingService()
}

func (sp *svcProvider) CallbackBuilding() isvc.CallbackBuildingService {
	return callback.NewCallbackBuildingService()
}

func (sp *svcProvider) DateTime() isvc.DateTimeService {
	return service.NewDateTimeService()
}

func (sp *svcProvider) User() isvc.UserService {
	return service.NewUserService(sp.repoProvider.User())
}

func (sp *svcProvider) SlotLocking() isvc.SlotLocking {
	slotLock, err := service.NewSlotLocking(sp.cacheProvider.RedisClient(), sp.cacheProvider.TTL())
	if err != nil {
		panic(err)
	}
	return slotLock
}

func (sp *svcProvider) Phone() isvc.PhoneService {
	return service.NewPhoneNormalizingService()
}

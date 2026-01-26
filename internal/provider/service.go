package provider

import (
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interface/provider"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interface/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service"
)

type svcProvider struct {
	repoProvider     iprovider.RepoProvider
	cacheProvider    iprovider.CacheProvider
	callbackProvider iprovider.CallbackProvider
}

func NewServiceProvider(
	repoProvider iprovider.RepoProvider,
	cacheProvider iprovider.CacheProvider,
	callbackProvider iprovider.CallbackProvider,
) iprovider.ServiceProvider {
	return &svcProvider{repoProvider: repoProvider, cacheProvider: cacheProvider, callbackProvider: callbackProvider}
}

func (sp *svcProvider) Slot() isvc.SlotService {
	return service.NewSlotService(sp.repoProvider.Slot(), sp.SlotLocker())
}

func (sp *svcProvider) Lock() isvc.LockService {
	return service.NewLockService(sp.SlotLocker(), sp.cacheProvider.SlotLock())
}

func (sp *svcProvider) Booking() isvc.BookingService {
	return service.NewBookingService(sp.repoProvider, sp.Slot(), sp.Price(), sp.DateTime())
}

func (sp *svcProvider) Price() isvc.PriceService {
	return service.NewPriceService(sp.repoProvider.Price())
}

func (sp *svcProvider) Config() isvc.ConfigService {
	return service.NewConfigService(sp.repoProvider.Config())
}

func (sp *svcProvider) DateTime() isvc.DateTimeService {
	return service.NewDateTimeService()
}

func (sp *svcProvider) User() isvc.UserService {
	return service.NewUserService(sp.repoProvider.User())
}

func (sp *svcProvider) SlotLocker() isvc.SlotLocker {
	slotLock, err := service.NewSlotLockerService(sp.cacheProvider.RedisClient(), sp.cacheProvider.TTL())
	if err != nil {
		panic(err)
	}
	return slotLock
}

func (sp *svcProvider) Phone() isvc.PhoneService {
	return service.NewPhoneNormalizingService()
}

func (sp *svcProvider) Statistics() isvc.StatisticService {
	return service.NewStatisticService()
}

package provider

import (
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service"
)

type svcProvider struct {
	repo     iprovider.RepoProvider
	cache    iprovider.CacheProvider
	callback iprovider.CallbackProvider
}

func NewServiceProvider(
	repo iprovider.RepoProvider,
	cache iprovider.CacheProvider,
	callback iprovider.CallbackProvider,
) iprovider.ServiceProvider {
	return &svcProvider{repo: repo, cache: cache, callback: callback}
}

func (sp *svcProvider) Slot() isvc.SlotService {
	return service.NewSlotService(sp.repo.Slot(), sp.SlotLocker())
}

func (sp *svcProvider) Lock() isvc.LockService {
	return service.NewLockService(sp.SlotLocker(), sp.cache.SlotLock())
}

func (sp *svcProvider) Booking() isvc.BookingService {
	return service.NewBookingService(sp.repo, sp.Slot(), sp.Price(), sp.DateTime())
}

func (sp *svcProvider) Price() isvc.PriceService {
	return service.NewPriceService(sp.repo.Price())
}

func (sp *svcProvider) Config() isvc.ConfigService {
	return service.NewConfigService(sp.repo.Config())
}

func (sp *svcProvider) DateTime() isvc.DateTimeService {
	return service.NewDateTimeService()
}

func (sp *svcProvider) User() isvc.UserService {
	return service.NewUserService(sp.repo.User())
}

func (sp *svcProvider) SlotLocker() isvc.SlotLocker {
	slotLock, err := service.NewSlotLockerService(sp.cache.RedisClient(), sp.cache.TTL())
	if err != nil {
		panic(err)
	}
	return slotLock
}

func (sp *svcProvider) Phone() isvc.PhoneService {
	return service.NewPhoneNormalizingService()
}

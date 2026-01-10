package service

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/cache"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	pg "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
	msgfmt "github.com/pan-asovsky/brandd-tg-bot/internal/service/message_formatting"
	"github.com/redis/go-redis/v9"
)

type Provider struct {
	pgProvider    *pg.Provider
	cacheProvider *cache.Provider
	tgapi         *api.BotAPI
	redisClient   *redis.Client
}

func NewProvider(pgProvider *pg.Provider, cacheProvider *cache.Provider, tgapi *api.BotAPI) *Provider {
	return &Provider{pgProvider: pgProvider, cacheProvider: cacheProvider, tgapi: tgapi}
}

func (p *Provider) Slot() i.SlotService {
	return &slotService{p.pgProvider.Slot(), p.SlotLocking()}
}

func (p *Provider) Keyboard() i.KeyboardService {
	return &keyboardService{callbackBuilding: p.CallbackBuilding(), dateTime: p.DateTime()}
}

func (p *Provider) Lock() i.LockService {
	return &lockService{p.SlotLocking(), p.cacheProvider.SlotLock()}
}

func (p *Provider) Booking() i.BookingService {
	return &bookingService{p.pgProvider, p.Slot(), p.Price()}
}

func (p *Provider) Telegram() i.TelegramService {
	return &telegramService{p.Keyboard(), p.DateTime(), p.MessageFormattingProvider(), p.tgapi}
}

func (p *Provider) Price() i.PriceService {
	return &priceService{p.pgProvider}
}

func (p *Provider) Config() i.ConfigService {
	return &configService{p.pgProvider.Config()}
}

func (p *Provider) CallbackParsing() i.CallbackParsingService {
	return &callbackParsingService{}
}

func (p *Provider) CallbackBuilding() i.CallbackBuildingService {
	return &callbackBuildingService{}
}

func (p *Provider) MessageFormattingProvider() msgfmt.MessageFormattingProviderService {
	return msgfmt.MessageFormattingProviderService{DateTime: p.DateTime()}
}

func (p *Provider) DateTime() i.DateTimeService {
	return &dateTimeService{}
}

func (p *Provider) User() i.UserService {
	return &userService{p.pgProvider.User()}
}

func (p *Provider) SlotLocking() i.SlotLocking {
	slotLock, err := NewSlotLocking(p.cacheProvider.RedisClient(), p.cacheProvider.TTL())
	if err != nil {
		panic(err)
	}
	return slotLock
}

func (p *Provider) Phone() i.PhoneService {
	return NewPhoneNormalizingService()
}

package service

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/cache"
	"github.com/pan-asovsky/brandd-tg-bot/internal/cache/locker"
	pg "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
)

type Provider struct {
	pgProvider *pg.Provider
	slotLocker *locker.SlotLocker
	lockCache  *cache.LockCache
	tgapi      *api.BotAPI
}

func NewProvider(pgProvider *pg.Provider, slotLocker *locker.SlotLocker, lockCache *cache.LockCache, tgapi *api.BotAPI) *Provider {
	return &Provider{pgProvider: pgProvider, slotLocker: slotLocker, lockCache: lockCache, tgapi: tgapi}
}

func (p *Provider) Slot() SlotService {
	return &slotService{p.pgProvider.Slot(), p.slotLocker}
}

func (p *Provider) Keyboard() KeyboardService {
	return &keyboardService{callbackService: p.BuildCallback()}
}

func (p *Provider) Lock() LockService {
	return &lockService{p.slotLocker, p.lockCache}
}

func (p *Provider) Booking() BookingService {
	return &bookingService{p.pgProvider, p.Slot(), p.Price()}
}

func (p *Provider) Telegram() TelegramService {
	return &telegramService{kb: p.Keyboard(), tgapi: p.tgapi}
}

func (p *Provider) Price() PriceService {
	return &priceService{pgProvider: p.pgProvider}
}

func (p *Provider) Config() ConfigService {
	return &configService{configRepo: p.pgProvider.Config()}
}

func (p *Provider) ParseCallback() ParseCallbackService {
	return &parseCallbackService{}
}

func (p *Provider) BuildCallback() BuildCallbackService {
	return BuildCallbackService{}
}

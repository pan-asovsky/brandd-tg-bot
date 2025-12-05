package service

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/cache"
	"github.com/pan-asovsky/brandd-tg-bot/internal/cache/locker"
	pg "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
)

type Provider struct {
	repoProvider *pg.Provider
	slotLocker   *locker.SlotLocker
	lockCache    *cache.LockCache
	tgapi        *api.BotAPI
}

func NewProvider(repoProvider *pg.Provider, slotLocker *locker.SlotLocker, lockCache *cache.LockCache, tgapi *api.BotAPI) *Provider {
	return &Provider{repoProvider: repoProvider, slotLocker: slotLocker, lockCache: lockCache, tgapi: tgapi}
}

func (p *Provider) Slot() SlotService {
	return &slotService{p.repoProvider.Slot(), p.slotLocker}
}

func (p *Provider) Keyboard() KeyboardService {
	return &keyboardService{}
}

func (p *Provider) Lock() LockService {
	return &lockService{p.slotLocker, p.lockCache}
}

func (p *Provider) Booking() BookingService {
	return &bookingService{p.repoProvider}
}

func (p *Provider) Telegram() TelegramService {
	return &telegramService{kb: p.Keyboard(), tgapi: p.tgapi}
}

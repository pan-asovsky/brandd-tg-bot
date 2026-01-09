package service

import (
	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/cache"
	"github.com/pan-asovsky/brandd-tg-bot/internal/cache/locker"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces"
	pg "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
	msgfmt "github.com/pan-asovsky/brandd-tg-bot/internal/service/message_formatting"
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

func (p *Provider) Slot() i.SlotService {
	return &slotService{p.pgProvider.Slot(), p.slotLocker}
}

func (p *Provider) Keyboard() i.KeyboardService {
	return &keyboardService{callbackBuilding: p.CallbackBuilding(), dateTime: p.DateTime()}
}

func (p *Provider) Lock() i.LockService {
	return &lockService{p.slotLocker, p.lockCache}
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

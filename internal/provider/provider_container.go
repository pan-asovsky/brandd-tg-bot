package provider

import iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"

type Container struct {
	Repo         iprovider.RepoProvider
	Service      iprovider.ServiceProvider
	Cache        iprovider.CacheProvider
	Telegram     iprovider.TelegramProvider
	Callback     iprovider.CallbackProvider
	Formatter    iprovider.MessageFormatterProvider
	Keyboard     iprovider.KeyboardProvider
	Notification iprovider.NotificationProvider
	Statistics   iprovider.StatisticsProvider
}

func NewContainer(
	repo iprovider.RepoProvider,
	service iprovider.ServiceProvider,
	cache iprovider.CacheProvider,
	telegram iprovider.TelegramProvider,
	callback iprovider.CallbackProvider,
	fmt iprovider.MessageFormatterProvider,
	keyboard iprovider.KeyboardProvider,
	notification iprovider.NotificationProvider,
	statistics iprovider.StatisticsProvider,
) *Container {
	return &Container{
		Repo:         repo,
		Service:      service,
		Cache:        cache,
		Telegram:     telegram,
		Callback:     callback,
		Formatter:    fmt,
		Keyboard:     keyboard,
		Notification: notification,
		Statistics:   statistics,
	}
}

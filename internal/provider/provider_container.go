package provider

import iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interface/provider"

type Container struct {
	RepoProvider         iprovider.RepoProvider
	ServiceProvider      iprovider.ServiceProvider
	CacheProvider        iprovider.CacheProvider
	TelegramProvider     iprovider.TelegramProvider
	CallbackProvider     iprovider.CallbackProvider
	MsgFmtProvider       iprovider.MessageFormatterProvider
	KeyboardProvider     iprovider.KeyboardProvider
	NotificationProvider iprovider.NotificationProvider
}

func NewContainer(
	repo iprovider.RepoProvider,
	service iprovider.ServiceProvider,
	cache iprovider.CacheProvider,
	telegram iprovider.TelegramProvider,
	callback iprovider.CallbackProvider,
	msgFmt iprovider.MessageFormatterProvider,
	keyboard iprovider.KeyboardProvider,
	notification iprovider.NotificationProvider,
) *Container {
	return &Container{
		RepoProvider:         repo,
		ServiceProvider:      service,
		CacheProvider:        cache,
		TelegramProvider:     telegram,
		CallbackProvider:     callback,
		MsgFmtProvider:       msgFmt,
		KeyboardProvider:     keyboard,
		NotificationProvider: notification,
	}
}

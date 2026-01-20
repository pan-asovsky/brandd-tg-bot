package provider

import iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"

type Container struct {
	RepoProvider     iprovider.RepoProvider
	ServiceProvider  iprovider.ServiceProvider
	CacheProvider    iprovider.CacheProvider
	TelegramProvider iprovider.TelegramProvider
	CallbackProvider iprovider.CallbackProvider
	MsgFmtProvider   iprovider.MessageFormatterProvider
	KeyboardProvider iprovider.KeyboardProvider
}

func NewContainer(
	repo iprovider.RepoProvider,
	service iprovider.ServiceProvider,
	cache iprovider.CacheProvider,
	telegram iprovider.TelegramProvider,
	callback iprovider.CallbackProvider,
	msgFmt iprovider.MessageFormatterProvider,
	keyboardProvider iprovider.KeyboardProvider,
) *Container {
	return &Container{
		RepoProvider:     repo,
		ServiceProvider:  service,
		CacheProvider:    cache,
		TelegramProvider: telegram,
		CallbackProvider: callback,
		MsgFmtProvider:   msgFmt,
		KeyboardProvider: keyboardProvider,
	}
}

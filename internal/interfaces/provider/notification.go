package provider

import inotif "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/notification"

type NotificationProvider interface {
	Service() inotif.Service
}

package provider

import inotif "github.com/pan-asovsky/brandd-tg-bot/internal/interface/notification"

type NotificationProvider interface {
	Service() inotif.Service
}

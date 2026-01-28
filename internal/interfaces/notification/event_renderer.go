package notification

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/model/notification"
)

type EventRenderer interface {
	Render(e notification.Event) (string, error)
	Register(t notification.Type, f notification.Formatter)
}

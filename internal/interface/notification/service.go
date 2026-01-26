package notification

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/model/notification"
)

type Service interface {
	Notify(e notification.Event) error
}

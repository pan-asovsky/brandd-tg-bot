package notification

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/model/notification"
)

type RecipientResolver interface {
	Resolve(e notification.Event) ([]notification.Recipient, error)
}

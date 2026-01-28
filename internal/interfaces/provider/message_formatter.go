package provider

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/fmt"
)

type MessageFormatterProvider interface {
	Booking() fmt.UserMessageFormatterService
	Admin() fmt.AdminMessageFormatterService
}

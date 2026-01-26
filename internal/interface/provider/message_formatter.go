package provider

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/interface/service/fmt"
)

type MessageFormatterProvider interface {
	Booking() fmt.UserMessageFormatterService
	Admin() fmt.AdminMessageFormatterService
}

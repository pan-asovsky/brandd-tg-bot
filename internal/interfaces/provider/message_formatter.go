package provider

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/msg_fmt"
)

type MessageFormatterProvider interface {
	Booking() msg_fmt.BookingMessageFormatterService
	Admin() msg_fmt.AdminMessageFormatterService
}

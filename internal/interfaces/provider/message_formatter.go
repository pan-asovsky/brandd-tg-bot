package provider

import "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"

type MessageFormatterProvider interface {
	Booking() service.BookingMessageFormatterService
	Admin() service.AdminMessageFormatterService
}

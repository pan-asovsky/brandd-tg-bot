package provider

import "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"

type MessageFormattingProvider interface {
	Booking() service.BookingMessageFormattingService
	Admin() service.AdminMessageFormattingService
}

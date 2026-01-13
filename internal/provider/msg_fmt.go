package provider

import (
	p "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service/msg_fmt"
)

type messageFormattingProvider struct {
	dateTimeSvc i.DateTimeService
}

func NewMessageFormattingProvider(dateTimeSvc i.DateTimeService) p.MessageFormattingProvider {
	return &messageFormattingProvider{dateTimeSvc: dateTimeSvc}
}

func (m *messageFormattingProvider) Booking() i.BookingMessageFormattingService {
	return msg_fmt.NewBookingMessageFormattingService(m.dateTimeSvc)
}

func (m *messageFormattingProvider) Admin() i.AdminMessageFormattingService {
	return msg_fmt.NewAdminMessageFormattingService(m.dateTimeSvc)
}

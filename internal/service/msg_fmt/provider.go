package msg_fmt

import (
	p "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
)

type messageFormattingProvider struct {
	dateTimeSvc i.DateTimeService
}

func NewMessageFormattingProvider(dateTimeSvc i.DateTimeService) p.MessageFormattingProvider {
	return &messageFormattingProvider{dateTimeSvc: dateTimeSvc}
}

func (m *messageFormattingProvider) Booking() i.BookingMessageFormattingService {
	return &bookingMessageFormattingService{dateTime: m.dateTimeSvc}
}

func (m *messageFormattingProvider) Admin() i.AdminMessageFormattingService {
	return &adminMessageFormattingService{dateTime: m.dateTimeSvc}
}

package message_formatting

import (
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
)

type MessageFormattingProviderService struct {
	DateTime i.DateTimeService
}

func (m *MessageFormattingProviderService) Booking() i.BookingMessageFormattingService {
	return &bookingMessageFormattingService{dateTime: m.DateTime}
}

func (m *MessageFormattingProviderService) Admin() i.AdminMessageFormattingService {
	return &adminMessageFormattingService{dateTime: m.DateTime}
}

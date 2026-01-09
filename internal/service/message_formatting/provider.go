package message_formatting

import i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces"

type MessageFormattingProviderService struct {
	DateTime i.DateTimeService
}

func (m *MessageFormattingProviderService) Booking() i.BookingMessageFormattingService {
	return &bookingMessageFormattingService{dateTime: m.DateTime}
}

package provider

import (
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service/msg_fmt"
)

type messageFormatterProvider struct {
	dateTimeSvc isvc.DateTimeService
}

func NewMessageFormatterProvider(dateTimeSvc isvc.DateTimeService) iprovider.MessageFormatterProvider {
	return &messageFormatterProvider{dateTimeSvc: dateTimeSvc}
}

func (m *messageFormatterProvider) Booking() isvc.BookingMessageFormatterService {
	return msg_fmt.NewBookingMessageFormattingService(m.dateTimeSvc)
}

func (m *messageFormatterProvider) Admin() isvc.AdminMessageFormatterService {
	return msg_fmt.NewAdminMessageFormattingService(m.dateTimeSvc)
}

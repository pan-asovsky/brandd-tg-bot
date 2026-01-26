package provider

import (
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interface/provider"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interface/service"
	ifmt "github.com/pan-asovsky/brandd-tg-bot/internal/interface/service/fmt"
	"github.com/pan-asovsky/brandd-tg-bot/internal/service/msg_fmt"
)

type messageFormatterProvider struct {
	dateTimeSvc isvc.DateTimeService
}

func NewMessageFormatterProvider(dateTimeSvc isvc.DateTimeService) iprovider.MessageFormatterProvider {
	return &messageFormatterProvider{dateTimeSvc: dateTimeSvc}
}

func (m *messageFormatterProvider) Booking() ifmt.UserMessageFormatterService {
	return msg_fmt.NewUserMessageFormatterService(m.dateTimeSvc)
}

func (m *messageFormatterProvider) Admin() ifmt.AdminMessageFormatterService {
	return msg_fmt.NewAdminMessageFormatterService(m.dateTimeSvc)
}

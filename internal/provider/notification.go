package provider

import (
	"fmt"

	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	inotif "github.com/pan-asovsky/brandd-tg-bot/internal/interface/notification"
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interface/provider"
	"github.com/pan-asovsky/brandd-tg-bot/internal/interface/service"
	itg "github.com/pan-asovsky/brandd-tg-bot/internal/interface/service/telegram"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model/notification"
	notif "github.com/pan-asovsky/brandd-tg-bot/internal/notification"
)

type notificationProvider struct {
	userService service.UserService
	tgCommon    itg.TelegramCommonService
	mFmt        iprovider.MessageFormatterProvider
}

func NewNotificationProvider(
	us service.UserService,
	tc itg.TelegramCommonService,
	mf iprovider.MessageFormatterProvider,
) iprovider.NotificationProvider {
	return &notificationProvider{userService: us, tgCommon: tc, mFmt: mf}
}

func (np *notificationProvider) Service() inotif.Service {
	return notif.NewNotificationService(
		notif.NewRecipientResolver(np.userService),
		np.registerFormatters(),
		[]inotif.Channel{
			notification.NewTelegramChannel(np.tgCommon),
		},
	)
}

func (np *notificationProvider) registerFormatters() inotif.EventRenderer {
	er := notif.NewEventRenderer()

	er.Register(notification.BookingCreated, func(data any) (string, error) {
		booking, ok := data.(*entity.Booking)
		if !ok {
			return "", fmt.Errorf("[notification] invalid payload for event %s", notification.BookingCreated)
		}
		return np.mFmt.Admin().BookingCreated(booking)
	})

	er.Register(notification.BookingCancelled, func(data any) (string, error) {
		booking, ok := data.(*entity.Booking)
		if !ok {
			return "", fmt.Errorf("[notification] invalid payload for event %s", notification.BookingCancelled)
		}
		return np.mFmt.Admin().BookingCancelled(booking)
	})

	er.Register(notification.BookingCompleted, func(data any) (string, error) {
		booking, ok := data.(*entity.Booking)
		if !ok {
			return "", fmt.Errorf("[notification] invalid payload for event %s", notification.BookingCompleted)
		}
		return np.mFmt.Admin().BookingCompleted(booking)
	})

	return er
}

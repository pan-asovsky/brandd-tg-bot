package notification

import (
	"fmt"

	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	inotif "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/notification"
	"github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model/notification"
)

type recipientResolver struct {
	adminsChatID []int64
}

func NewRecipientResolver(userService service.UserService) inotif.RecipientResolver {
	admins := userService.GetActiveAdmins()
	adminsChatID := make([]int64, len(admins))
	for _, admin := range admins {
		adminsChatID = append(adminsChatID, admin.ChatID)
	}

	return &recipientResolver{adminsChatID: adminsChatID}
}

func (r *recipientResolver) Resolve(e notification.Event) ([]notification.Recipient, error) {
	switch e.Type {

	case notification.BookingCreated, notification.BookingCancelled:
		rec := make([]notification.Recipient, 0, len(r.adminsChatID))
		for _, chatID := range r.adminsChatID {
			rec = append(rec, notification.Recipient{ChatID: chatID})
		}
		return rec, nil

	case notification.BookingCompleted:
		booking, ok := e.Data.(*entity.Booking)
		if !ok {
			return nil, fmt.Errorf("[resolve_recipients] for event %s received invalid payload %v", e.Data, e.Type)
		}
		return []notification.Recipient{
			{ChatID: booking.ChatID},
		}, nil
	}

	return nil, nil
}

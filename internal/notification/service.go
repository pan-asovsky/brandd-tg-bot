package notification

import (
	"log"

	inotif "github.com/pan-asovsky/brandd-tg-bot/internal/interface/notification"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model/notification"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type notificationService struct {
	recipients  inotif.RecipientResolver
	eventRender inotif.EventRenderer
	channels    []inotif.Channel
}

func NewNotificationService(
	rr inotif.RecipientResolver,
	er inotif.EventRenderer,
	ch []inotif.Channel,
) inotif.Service {
	return &notificationService{recipients: rr, eventRender: er, channels: ch}
}

func (ns *notificationService) Notify(e notification.Event) error {
	recipients, err := ns.recipients.Resolve(e)
	if err != nil {
		return utils.WrapError(err)
	}

	msg, err := ns.eventRender.Render(e)
	if err != nil {
		return utils.WrapError(err)
	}

	for _, r := range recipients {
		for _, ch := range ns.channels {
			if err = ch.Send(r.ChatID, msg); err != nil {
				log.Printf("[notification_service] an error ocurred: %v", err)
			}
		}
	}

	return nil
}

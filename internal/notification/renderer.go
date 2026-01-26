package notification

import (
	"fmt"

	inotif "github.com/pan-asovsky/brandd-tg-bot/internal/interface/notification"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model/notification"
)

type eventRenderer struct {
	formatters map[notification.Type]notification.Formatter
}

func NewEventRenderer() inotif.EventRenderer {
	return &eventRenderer{
		formatters: make(map[notification.Type]notification.Formatter),
	}
}

func (er *eventRenderer) Register(t notification.Type, f notification.Formatter) {
	er.formatters[t] = f
}

func (er *eventRenderer) Render(e notification.Event) (string, error) {
	f, ok := er.formatters[e.Type]
	if !ok {
		return "", fmt.Errorf("[render_event] no formatter for type: %s", e.Type)
	}
	return f(e.Data)
}

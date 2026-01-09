package service

import (
	"errors"
	"log"
	"strings"

	api "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
)

type callbackParsingService struct {
}

const (
	KeyDate    = "D"
	KeyZone    = "Z"
	KeyTime    = "T"
	KeyService = "S"
	KeyRadius  = "R"
)

func (c *callbackParsingService) Parse(query *api.CallbackQuery) (*types.UserSessionInfo, error) {
	log.Printf("[parse_callback] callback: %s", query.Data)
	_, payload, ok := strings.Cut(query.Data, "::")
	if !ok {
		return nil, errors.New("[parse_callback] split callback error " + query.Data)
	}

	var info types.UserSessionInfo

	parts := strings.Split(payload, "|")
	for _, part := range parts {
		key, value, ok := strings.Cut(part, "~")
		if !ok {
			return nil, errors.New("[parse_callback] callback error: " + payload)
		}

		switch key {
		case KeyDate:
			info.Date = decodeDate(value)
		case KeyZone:
			info.Zone = decodeTime(value)
		case KeyTime:
			info.Time = decodeTime(value)
		case KeyService:
			info.Service = value
		case KeyRadius:
			info.RimRadius = value
		default:
			log.Printf("Unknown key: %s", key)

		}
	}

	info.ChatID = query.Message.Chat.ID
	return &info, nil
}

func decodeDate(encoded string) string {
	if len(encoded) == 8 {
		return encoded[:4] + "-" + encoded[4:6] + "-" + encoded[6:]
	}
	return encoded
}

func decodeTime(encoded string) string {
	if len(encoded) == 4 {
		return encoded[:2] + ":00-" + encoded[2:] + ":00"
	}
	return encoded
}

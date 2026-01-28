package callback

import (
	"errors"
	"fmt"
	"strings"

	tgapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constant/admin_flow"
	icallback "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/callback"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type adminCallbackParserService struct{}

func NewAdminCallbackParserService() icallback.AdminCallbackParserService {
	return &adminCallbackParserService{}
}

func (acps *adminCallbackParserService) ParseNoShow(query *tgapi.CallbackQuery) (*model.BookingInfo, error) {
	info, ok := strings.CutPrefix(query.Data, admflow.PrefixNoShow)
	if !ok {
		return nil, errors.New("[parse_no_show] prefix NS not found: " + query.Data)
	}

	details := strings.Split(info, ":")
	if len(details) != 3 {
		return nil, errors.New(fmt.Sprintf("[parse_no_show] size: %d. info: %s", len(details), info))
	}

	return model.NewNoShowBookingInfo(details)
}

func (acps *adminCallbackParserService) ParseComplete(query *tgapi.CallbackQuery) (*model.BookingInfo, error) {
	info, ok := strings.CutPrefix(query.Data, admflow.PrefixComplete)
	if !ok {
		return nil, errors.New("[parse_complete] prefix CMP not found: " + query.Data)
	}

	details := strings.Split(info, ":")
	if len(details) != 3 {
		return nil, errors.New(fmt.Sprintf("[parse_complete] size: %d. info: %s", len(details), info))
	}

	return model.NewCompleteBookingInfo(details)
}

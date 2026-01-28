package callback

import (
	"fmt"
	"strconv"

	admflow "github.com/pan-asovsky/brandd-tg-bot/internal/constant/admin_flow"
	icallback "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/callback"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type adminCallbackBuilderService struct{}

func NewAdminCallbackBuilderService() icallback.AdminCallbackBuilderService {
	return &adminCallbackBuilderService{}
}

func (acbs *adminCallbackBuilderService) StartAdmin() string {
	return admflow.AdminFlowCbk
}

func (acbs *adminCallbackBuilderService) StartUser() string {
	return admflow.UserFlowCbk
}

func (acbs *adminCallbackBuilderService) BookingsMenu() string {
	return admflow.BookingsCbk
}

func (acbs *adminCallbackBuilderService) Booking(bookingID int64) string {
	return admflow.AdminPrefix + admflow.PrefixBooking + strconv.FormatInt(bookingID, 10)
}

func (acbs *adminCallbackBuilderService) Statistics() string {
	return admflow.StatisticsCbk
}

func (acbs *adminCallbackBuilderService) Settings() string {
	return admflow.SettingsCbk
}

func (acbs *adminCallbackBuilderService) Back(direction string) string {
	return admflow.AdminBackPrefix + direction
}

func (acbs *adminCallbackBuilderService) Chat(userChatID int64) string {
	return fmt.Sprintf("tg://user?id=%d", userChatID)
}

func (acbs *adminCallbackBuilderService) PreComplete(userChatID int64, bookingID int64) string {
	return fmt.Sprintf(
		"%s%s%d:%d",
		admflow.AdminPrefix,
		admflow.PrefixPreCompleteBooking,
		userChatID,
		bookingID,
	)
}

func (acbs *adminCallbackBuilderService) PreNoShow(userChatID int64, bookingID int64) string {
	return fmt.Sprintf(
		"%s%s%d:%d",
		admflow.AdminPrefix,
		admflow.PrefixPreNoShowBooking,
		userChatID,
		bookingID,
	)
}

func (acbs *adminCallbackBuilderService) Confirm(info *model.BookingInfo) string {
	var statusPart string
	switch info.Status {
	case model.PreCompleted:
		statusPart = admflow.PrefixCompleteBooking
	case model.PreNoShow:
		statusPart = admflow.PrefixNoShowBooking
	}

	return fmt.Sprintf(
		"%s%s%d:%d",
		admflow.AdminPrefix,
		statusPart,
		info.UserChatID,
		info.BookingID,
	)
}

func (acbs *adminCallbackBuilderService) Reject(info *model.BookingInfo) string {
	return admflow.RejectActionCbk + string(info.Status)
}

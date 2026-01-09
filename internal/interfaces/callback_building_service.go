package interfaces

import (
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
)

type CallbackBuildingService interface {
	Menu() string
	NewBooking() string
	MyBookings() string
	PreCancelBooking() string
	CancelBooking() string
	NoCancelBooking() string
	Date(date time.Time) string
	Zone(date, zone string) string
	Time(info *types.UserSessionInfo) string
	ServiceSelection(service string, info *types.UserSessionInfo) string
	ServiceConfirmation(info *types.UserSessionInfo) string
	Rim(info *types.UserSessionInfo) string
}

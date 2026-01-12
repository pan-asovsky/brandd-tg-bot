package service

import (
	"time"

	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
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
	Time(info *model.UserSessionInfo) string
	ServiceSelection(service string, info *model.UserSessionInfo) string
	ServiceConfirmation(info *model.UserSessionInfo) string
	Rim(info *model.UserSessionInfo) string

	//todo: split service to user and admin (+ maybe common)
	StartAdmin() string
	StartUser() string
}

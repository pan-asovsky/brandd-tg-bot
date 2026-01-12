package service

import (
	"fmt"
	"strings"
	"time"

	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

const (
	Date    = "D~"
	Zone    = "Z~"
	Time    = "T~"
	Service = "S~"
	Rim     = "R~"
)

type callbackBuildingService struct{}

func (cbs *callbackBuildingService) Menu() string {
	return usflow.UserPrefix + usflow.PrefixBack + usflow.Menu
}

func (cbs *callbackBuildingService) NewBooking() string {
	return usflow.NewBookingCbk
}

func (cbs *callbackBuildingService) MyBookings() string {
	return usflow.MyBookingsCbk
}

func (cbs *callbackBuildingService) PreCancelBooking() string {
	return usflow.PreCancelBookingCbk
}

func (cbs *callbackBuildingService) CancelBooking() string {
	return usflow.CancelBookingCbk
}

func (cbs *callbackBuildingService) NoCancelBooking() string {
	return usflow.NoCancelBookingCbk
}

func (cbs *callbackBuildingService) Date(date time.Time) string {
	return fmt.Sprintf("%s%s%s",
		usflow.UserPrefix+usflow.PrefixDate,
		Date, encodeDate(date.Format("2006-01-02")),
	)
}

func (cbs *callbackBuildingService) Zone(date, zone string) string {
	return fmt.Sprintf("%s%s%s|%s%s",
		usflow.UserPrefix+usflow.PrefixZone,
		Date, encodeDate(date),
		Zone, encodeTime(zone),
	)
}

func (cbs *callbackBuildingService) Time(info *model.UserSessionInfo) string {
	return fmt.Sprintf("%s%s%s|%s%s|%s%s",
		usflow.UserPrefix+usflow.PrefixTime,
		Date, encodeDate(info.Date),
		Zone, encodeTime(info.Zone),
		Time, encodeTime(info.Time),
	)
}

func (cbs *callbackBuildingService) ServiceSelection(service string, info *model.UserSessionInfo) string {
	return fmt.Sprintf("%s%s%s|%s%s|%s%s|%s%s",
		usflow.UserPrefix+usflow.PrefixServiceSelect,
		Date, encodeDate(info.Date),
		Time, encodeTime(info.Time),
		Zone, encodeTime(info.Zone),
		Service, service,
	)
}

func (cbs *callbackBuildingService) ServiceConfirmation(info *model.UserSessionInfo) string {
	return fmt.Sprintf("%s%s%s|%s%s|%s%s|%s%s",
		usflow.UserPrefix+usflow.PrefixServiceConfirm,
		Date, encodeDate(info.Date),
		Zone, encodeTime(info.Zone),
		Time, encodeTime(info.Time),
		Service, info.Service,
	)
}

func (cbs *callbackBuildingService) Rim(info *model.UserSessionInfo) string {
	return fmt.Sprintf("%s%s%s|%s%s|%s%s|%s%s|%s%s",
		usflow.UserPrefix+usflow.PrefixRim,
		Date, encodeDate(info.Date),
		Zone, encodeTime(info.Zone),
		Time, encodeTime(info.Time),
		Service, info.Service,
		Rim, info.RimRadius,
	)
}

func (cbs *callbackBuildingService) StartAdmin() string {
	return "FLOW::ADMIN"
}

func (cbs *callbackBuildingService) StartUser() string {
	return "FLOW::USER"
}

func encodeDate(date string) string {
	return strings.ReplaceAll(date, "-", "")
}

func encodeTime(zone string) string {
	withoutZeros := strings.NewReplacer(":00", "").Replace(zone)
	return strings.ReplaceAll(withoutZeros, "-", "")
}

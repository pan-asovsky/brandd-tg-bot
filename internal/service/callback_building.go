package service

import (
	"fmt"
	"strings"
	"time"

	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
)

const (
	Date    = "D~"
	Zone    = "Z~"
	Time    = "T~"
	Service = "S~"
	Rim     = "R~"
)

type callbackBuildingService struct{}

func (cb *callbackBuildingService) Menu() string {
	return consts.PrefixBack + consts.Menu
}

func (cb *callbackBuildingService) NewBooking() string {
	return consts.NewBookingCbk
}

func (cb *callbackBuildingService) MyBookings() string {
	return consts.MyBookingsCbk
}

func (cb *callbackBuildingService) PreCancelBooking() string {
	return consts.PreCancelBookingCbk
}

func (cb *callbackBuildingService) CancelBooking() string {
	return consts.CancelBookingCbk
}

func (cb *callbackBuildingService) NoCancelBooking() string {
	return consts.NoCancelBookingCbk
}

func (cb *callbackBuildingService) Date(date time.Time) string {
	return fmt.Sprintf("%s%s%s",
		consts.PrefixDate,
		Date, encodeDate(date.Format("2006-01-02")),
	)
}

func (cb *callbackBuildingService) Zone(date, zone string) string {
	return fmt.Sprintf("%s%s%s|%s%s",
		consts.PrefixZone,
		Date, encodeDate(date),
		Zone, encodeTime(zone),
	)
}

func (cb *callbackBuildingService) Time(info *types.UserSessionInfo) string {
	return fmt.Sprintf("%s%s%s|%s%s|%s%s",
		consts.PrefixTime,
		Date, encodeDate(info.Date),
		Zone, encodeTime(info.Zone),
		Time, encodeTime(info.Time),
	)
}

func (cb *callbackBuildingService) ServiceSelection(service string, info *types.UserSessionInfo) string {
	return fmt.Sprintf("%s%s%s|%s%s|%s%s|%s%s",
		consts.PrefixServiceSelect,
		Date, encodeDate(info.Date),
		Time, encodeTime(info.Time),
		Zone, encodeTime(info.Zone),
		Service, service,
	)
}

func (cb *callbackBuildingService) ServiceConfirmation(info *types.UserSessionInfo) string {
	return fmt.Sprintf("%s%s%s|%s%s|%s%s|%s%s",
		consts.PrefixServiceConfirm,
		Date, encodeDate(info.Date),
		Zone, encodeTime(info.Zone),
		Time, encodeTime(info.Time),
		Service, info.Service,
	)
}

func (cb *callbackBuildingService) Rim(info *types.UserSessionInfo) string {
	return fmt.Sprintf("%s%s%s|%s%s|%s%s|%s%s|%s%s",
		consts.PrefixRim,
		Date, encodeDate(info.Date),
		Zone, encodeTime(info.Zone),
		Time, encodeTime(info.Time),
		Service, info.Service,
		Rim, info.RimRadius,
	)
}

func encodeDate(date string) string {
	return strings.ReplaceAll(date, "-", "")
}

func encodeTime(zone string) string {
	withoutZeros := strings.NewReplacer(":00", "").Replace(zone)
	return strings.ReplaceAll(withoutZeros, "-", "")
}

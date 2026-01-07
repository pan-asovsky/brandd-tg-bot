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

type BuildCallbackService struct{}

func (BuildCallbackService) Menu() string {
	return consts.PrefixBack + consts.Menu
}

func (BuildCallbackService) NewBooking() string {
	return consts.NewBookingCbk
}

func (BuildCallbackService) MyBookings() string {
	return consts.MyBookingsCbk
}

func (BuildCallbackService) PreCancelBooking() string {
	return consts.PreCancelBookingCbk
}

func (BuildCallbackService) CancelBooking() string {
	return consts.CancelBookingCbk
}

func (BuildCallbackService) NoCancelBooking() string {
	return consts.NoCancelBookingCbk
}

func (BuildCallbackService) Date(date time.Time) string {
	return fmt.Sprintf("%s%s%s",
		consts.PrefixDate,
		Date, encodeDate(date.Format("2006-01-02")),
	)
}

func (BuildCallbackService) Zone(date, zone string) string {
	return fmt.Sprintf("%s%s%s|%s%s",
		consts.PrefixZone,
		Date, encodeDate(date),
		Zone, encodeTime(zone),
	)
}

func (BuildCallbackService) Time(info *types.UserSessionInfo) string {
	return fmt.Sprintf("%s%s%s|%s%s|%s%s",
		consts.PrefixTime,
		Date, encodeDate(info.Date),
		Zone, encodeTime(info.Zone),
		Time, encodeTime(info.Time),
	)
}

func (BuildCallbackService) ServiceSelection(service string, info *types.UserSessionInfo) string {
	return fmt.Sprintf("%s%s%s|%s%s|%s%s|%s%s",
		consts.PrefixServiceSelect,
		Date, encodeDate(info.Date),
		Time, encodeTime(info.Time),
		Zone, encodeTime(info.Zone),
		Service, service,
	)
}

func (BuildCallbackService) ServiceConfirmation(info *types.UserSessionInfo) string {
	return fmt.Sprintf("%s%s%s|%s%s|%s%s|%s%s",
		consts.PrefixServiceConfirm,
		Date, encodeDate(info.Date),
		Zone, encodeTime(info.Zone),
		Time, encodeTime(info.Time),
		Service, info.Service,
	)
}

func (BuildCallbackService) Rim(info *types.UserSessionInfo) string {
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

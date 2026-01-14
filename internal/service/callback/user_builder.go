package callback

import (
	"fmt"
	"strings"
	"time"

	usflow "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	icallback "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service/callback"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

const (
	Date    = "D~"
	Zone    = "Z~"
	Time    = "T~"
	Service = "S~"
	Rim     = "R~"
)

type userCallbackBuilderService struct{}

func NewUserCallbackBuilderService() icallback.UserCallbackBuilderService {
	return &userCallbackBuilderService{}
}

func (ucbs *userCallbackBuilderService) Menu() string {
	return usflow.UserPrefix + usflow.PrefixBack + usflow.Menu
}

func (ucbs *userCallbackBuilderService) NewBooking() string {
	return usflow.NewBookingCbk
}

func (ucbs *userCallbackBuilderService) MyBookings() string {
	return usflow.MyBookingsCbk
}

func (ucbs *userCallbackBuilderService) PreCancelBooking() string {
	return usflow.PreCancelBookingCbk
}

func (ucbs *userCallbackBuilderService) CancelBooking() string {
	return usflow.CancelBookingCbk
}

func (ucbs *userCallbackBuilderService) NoCancelBooking() string {
	return usflow.NoCancelBookingCbk
}

func (ucbs *userCallbackBuilderService) Date(date time.Time) string {
	return fmt.Sprintf("%s%s%s",
		usflow.UserPrefix+usflow.PrefixDate,
		Date, encodeDate(date.Format("2006-01-02")),
	)
}

func (ucbs *userCallbackBuilderService) Zone(date, zone string) string {
	return fmt.Sprintf("%s%s%s|%s%s",
		usflow.UserPrefix+usflow.PrefixZone,
		Date, encodeDate(date),
		Zone, encodeTime(zone),
	)
}

func (ucbs *userCallbackBuilderService) Time(info *model.UserSessionInfo) string {
	return fmt.Sprintf("%s%s%s|%s%s|%s%s",
		usflow.UserPrefix+usflow.PrefixTime,
		Date, encodeDate(info.Date),
		Zone, encodeTime(info.Zone),
		Time, encodeTime(info.Time),
	)
}

func (ucbs *userCallbackBuilderService) ServiceSelection(service string, info *model.UserSessionInfo) string {
	return fmt.Sprintf("%s%s%s|%s%s|%s%s|%s%s",
		usflow.UserPrefix+usflow.PrefixServiceSelect,
		Date, encodeDate(info.Date),
		Time, encodeTime(info.Time),
		Zone, encodeTime(info.Zone),
		Service, service,
	)
}

func (ucbs *userCallbackBuilderService) ServiceConfirmation(info *model.UserSessionInfo) string {
	return fmt.Sprintf("%s%s%s|%s%s|%s%s|%s%s",
		usflow.UserPrefix+usflow.PrefixServiceConfirm,
		Date, encodeDate(info.Date),
		Zone, encodeTime(info.Zone),
		Time, encodeTime(info.Time),
		Service, info.Service,
	)
}

func (ucbs *userCallbackBuilderService) Rim(info *model.UserSessionInfo) string {
	return fmt.Sprintf("%s%s%s|%s%s|%s%s|%s%s|%s%s",
		usflow.UserPrefix+usflow.PrefixRim,
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

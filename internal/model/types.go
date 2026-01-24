package model

import (
	"strconv"

	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type SlotLockInfo struct {
	Key  string
	UUID string
}

type UserSessionInfo struct {
	ChatID                               int64
	Date, Zone, Time, RimRadius, Service string
	SelectedServices                     map[string]bool
	TotalPrice                           int64
}

type BookingInfo struct {
	UserChatID int64
	BookingID  int64
	Status     Status
}

func NewNoShowBookingInfo(details []string) (*BookingInfo, error) {
	var status Status
	switch details[0] {
	case "1":
		status = PreNoShow
	case "2":
		status = NoShow
	}

	userChatID, err := strconv.ParseInt(details[1], 10, 64)
	if err != nil {
		return nil, utils.WrapError(err)
	}

	bookingID, err := strconv.ParseInt(details[2], 10, 64)
	if err != nil {
		return nil, utils.WrapError(err)
	}

	return &BookingInfo{UserChatID: userChatID, BookingID: bookingID, Status: status}, nil
}

type Status string

const (
	PreNoShow   Status = "PreNoShow"
	NoShow      Status = "NoShow"
	PreComplete Status = "PreComplete"
	Complete    Status = "Complete"
)

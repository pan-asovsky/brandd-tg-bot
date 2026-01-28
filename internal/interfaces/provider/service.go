package provider

import (
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
)

type ServiceProvider interface {
	Slot() isvc.SlotService
	Lock() isvc.LockService
	Booking() isvc.BookingService
	Price() isvc.PriceService
	Config() isvc.ConfigService
	DateTime() isvc.DateTimeService
	User() isvc.UserService
	SlotLocker() isvc.SlotLocker
	Phone() isvc.PhoneService
}

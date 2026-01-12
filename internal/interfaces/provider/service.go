package provider

import i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"

type ServiceProvider interface {
	Slot() i.SlotService
	Keyboard() i.KeyboardService
	Lock() i.LockService
	Booking() i.BookingService
	Price() i.PriceService
	Config() i.ConfigService
	CallbackParsing() i.CallbackParsingService
	CallbackBuilding() i.CallbackBuildingService
	DateTime() i.DateTimeService
	User() i.UserService
	SlotLocking() i.SlotLocking
	Phone() i.PhoneService
}

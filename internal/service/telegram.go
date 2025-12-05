package service

type TelegramService interface {
	ProcessMenu(bookings []AvailableBooking, chatID int64)
	ProcessDate()
	ProcessZone()
	ProcessTime()
	ProcessServiceType()
	ProcessRimRadius()
	ProcessPhone()
}

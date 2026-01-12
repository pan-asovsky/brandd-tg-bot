package provider

import i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"

type RepoProvider interface {
	Price() i.PriceRepo
	Config() i.ConfigRepo
	Slot() i.SlotRepo
	Booking() i.BookingRepo
	User() i.UserRepo
	Service() i.ServiceRepo
}

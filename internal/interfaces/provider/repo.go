package provider

import irepo "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"

type RepoProvider interface {
	Price() irepo.PriceRepo
	Config() irepo.ConfigRepo
	Slot() irepo.SlotRepo
	Booking() irepo.BookingRepo
	User() irepo.UserRepo
	Service() irepo.ServiceRepo
}

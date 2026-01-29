package statistics

import (
	"log"

	irepo "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model/stat"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type statisticService struct {
	bookingRepo irepo.BookingRepo
}

func NewStatisticService(br irepo.BookingRepo) isvc.StatisticService {
	return &statisticService{bookingRepo: br}
}

func (ss *statisticService) Calculate(p stat.Period) (stat.Stats, error) {
	log.Printf("[calculate_statistics] period %s: %s", p.Label, p.Format())

	bookings, err := ss.bookingRepo.ListByPeriod(p)
	if err != nil {
		return stat.Stats{}, utils.WrapError(err)
	}

	var stats stat.Stats
	for _, booking := range bookings {
		switch {
		case booking.IsActive():
			stats.ActiveCount++
		case booking.IsCancelled():
			stats.CanceledCount++
		case booking.IsCompleted():
			stats.CompletedCount++
		case booking.IsNoShow():
			stats.NoShowCount++
		case booking.IsPending():
			stats.PendingCount++
		}
	}

	return stats, nil
}

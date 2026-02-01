package statistics

import (
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
	return utils.WrapFunction(func() (stat.Stats, error) {
		return ss.bookingRepo.StatusesByPeriod(p)
	})
}

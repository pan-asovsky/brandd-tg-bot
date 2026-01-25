package service

import (
	"log"

	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
)

type statisticService struct {
}

func NewStatisticService() isvc.StatisticService {
	return &statisticService{}
}

func (s statisticService) Add(booking *entity.Booking) error {
	log.Printf("[add_statistic] booking: %v", booking)

	return nil
}

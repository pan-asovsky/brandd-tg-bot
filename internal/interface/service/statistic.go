package service

import (
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
)

type StatisticService interface {
	Add(booking *entity.Booking) error
}

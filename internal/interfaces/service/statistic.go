package service

import "github.com/pan-asovsky/brandd-tg-bot/internal/model/stat"

type StatisticService interface {
	Calculate(p stat.Period) (stat.Stats, error)
}

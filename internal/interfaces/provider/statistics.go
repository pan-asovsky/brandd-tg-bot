package provider

import (
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	stat "github.com/pan-asovsky/brandd-tg-bot/internal/statistics"
)

type StatisticsProvider interface {
	Service() isvc.StatisticService
	PeriodFactory() *stat.PeriodFactory
}

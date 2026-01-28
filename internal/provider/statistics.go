package provider

import (
	iprovider "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/provider"
	isvc "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	stat "github.com/pan-asovsky/brandd-tg-bot/internal/statistics"
)

type statisticsProvider struct {
	repo iprovider.RepoProvider
}

func NewStatisticsProvider(repo iprovider.RepoProvider) iprovider.StatisticsProvider {
	return &statisticsProvider{repo: repo}
}

func (sp *statisticsProvider) Service() isvc.StatisticService {
	return stat.NewStatisticService(sp.repo.Booking())
}

func (sp *statisticsProvider) PeriodFactory() *stat.PeriodFactory {
	return stat.NewPeriodFactory()
}

package mocks

import "github.com/pan-asovsky/brandd-tg-bot/internal/entity"

type SlotRepoMock struct {
	TodayAvailable bool
}

func (m *SlotRepoMock) IsTodayAvailable() bool {
	return m.TodayAvailable
}

func (m *SlotRepoMock) GetAvailableSlots(date string) ([]entity.Slot, error) {
	return []entity.Slot{}, nil
}

func (m *SlotRepoMock) FindByDateAndTime(date, time string) (*entity.Slot, error) {
	return &entity.Slot{}, nil
}

func (m *SlotRepoMock) MarkUnavailable(date, startTime string) error {
	return nil
}

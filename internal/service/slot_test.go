package service

import (
	"testing"
	"time"

	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants/user_flow"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func NewSlotService(todayAvailable bool) i.SlotService {
	return &slotService{slotRepo: &mocks.SlotRepoMock{TodayAvailable: todayAvailable}, slotLocker: mocks.NewSlotLockingMock()}
}

func TestGetAvailableBookings_TodayAvailable(t *testing.T) {
	slotSvc := NewSlotService(true)
	result := slotSvc.GetAvailableBookings()
	assert.Len(t, result, 3)

	now := time.Now()

	assert.Equal(t, consts.Today, result[0].Label)
	assert.True(t, sameDay(now, result[0].Date))

	assert.Equal(t, consts.Tomorrow, result[1].Label)
	assert.True(t, sameDay(now.AddDate(0, 0, 1), result[1].Date))

	assert.Equal(t, consts.AfterTomorrow, result[2].Label)
	assert.True(t, sameDay(now.AddDate(0, 0, 2), result[2].Date))
}

func TestGetAvailableBookings_TodayNotAvailable(t *testing.T) {
	slotSvc := NewSlotService(false)
	result := slotSvc.GetAvailableBookings()
	assert.Len(t, result, 2)

	now := time.Now()

	assert.Equal(t, consts.Tomorrow, result[0].Label)
	assert.True(t, sameDay(now.AddDate(0, 0, 1), result[0].Date))

	assert.Equal(t, consts.AfterTomorrow, result[1].Label)
	assert.True(t, sameDay(now.AddDate(0, 0, 2), result[1].Date))
}

func sameDay(a, b time.Time) bool {
	y1, m1, d1 := a.Date()
	y2, m2, d2 := b.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

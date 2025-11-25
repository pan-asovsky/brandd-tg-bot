package service

import (
	"testing"

	"github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type mockSlotRepo struct {
	isTodayAvailable bool
}

func (m *mockSlotRepo) IsTodayAvailable() bool {
	return m.isTodayAvailable
}

func (m *mockSlotRepo) GetAvailableSlots(date string) ([]model.Slot, error) {
	return nil, nil
}

func (m *mockSlotRepo) FindByZone(zone model.Zone) []model.Slot {
	return nil
}

func (m *mockSlotRepo) FindByTimeslot(timeslot model.Timeslot) []model.Slot {
	return nil
}

func TestGetAvailableBookings_TodayAvailable(t *testing.T) {
	repo := &mockSlotRepo{isTodayAvailable: true}
	svc := NewSlot(repo)

	bookings := svc.GetAvailableBookings()

	if len(bookings) != 3 {
		t.Fatalf("excepted 3 bookings, got %d", len(bookings))
	}
}

func TestGetAvailableBookings_TodayNotAvailable(t *testing.T) {
	repo := &mockSlotRepo{isTodayAvailable: false}
	svc := NewSlot(repo)

	bookings := svc.GetAvailableBookings()

	if len(bookings) != 2 {
		t.Fatalf("expected 2 bookings, got %d", len(bookings))
	}

	if bookings[0].Label != constants.Tomorrow {
		t.Errorf("expected first label %s, got %s", constants.Tomorrow, bookings[0].Label)
	}
}

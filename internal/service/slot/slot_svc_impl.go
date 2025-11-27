package service

import (
	"fmt"
	"log"
	"time"

	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	pg "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
	rd "github.com/pan-asovsky/brandd-tg-bot/internal/repository/redis"
)

type slotService struct {
	slotRepo pg.SlotRepository
	cache    rd.ZoneCache
}

func NewSlot(slotRepo pg.SlotRepository) SlotService {
	return &slotService{slotRepo: slotRepo}
}

func (s *slotService) GetAvailableBookings() []AvailableBooking {
	today := time.Now()

	days := []struct {
		offset    int
		label     string
		available bool
	}{
		{0, consts.Today, s.slotRepo.IsTodayAvailable()},
		{1, consts.Tomorrow, true},
		{2, consts.AfterTomorrow, true},
	}

	bookings := make([]AvailableBooking, 0, 3)
	for _, day := range days {
		if !day.available {
			continue
		}
		bookings = append(bookings, AvailableBooking{
			Date:  today.AddDate(0, 0, day.offset),
			Label: day.label,
		})
	}

	return bookings
}
func (s *slotService) GetAvailableZones(date string) (model.Zone, error) {
	slots, err := s.slotRepo.GetAvailableSlots(date)
	if err != nil {
		return nil, fmt.Errorf("error getting available slots: %w", err)
	}

	return s.groupByZones(slots), nil
}

func (s *slotService) groupByZones(slots []model.Slot) model.Zone {
	zones := make(model.Zone)

	for _, slot := range slots {
		if !slot.IsAvailable {
			continue
		}

		slotTime, err := time.Parse("15:04", slot.StartTime)
		if err != nil {
			log.Printf("[group_by_zones] error parsing slot time %s: %v", slot.StartTime, err)
			continue
		}

		for _, z := range model.ZonesDefinition {
			startTime, _ := time.Parse("15:04", z.Start)
			endTime, _ := time.Parse("15:04", z.End)

			if !slotTime.Before(startTime) && slotTime.Before(endTime) {
				zones[z.Name] = append(zones[z.Name], model.Timeslot{
					Start: slot.StartTime,
					End:   slot.EndTime,
				})
				break
			}
		}
	}
	return zones
}

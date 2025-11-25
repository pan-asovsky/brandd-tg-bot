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

	log.Printf("Retrieved %d slots for date %s", len(slots), date)

	return s.groupByZones(slots), nil
}

func (s *slotService) groupByZones(slots []model.Slot) model.Zone {
	zones := make(model.Zone)
	for _, slot := range slots {
		if !slot.IsAvailable {
			continue
		}

		slotHour := time.Date(0, 1, 1, slot.StartTime.Hour(), slot.StartTime.Minute(), 0, 0, time.UTC)
		for _, z := range model.ZonesDefinition {
			startTime, _ := time.Parse("15:04", z.Start)
			endTime, _ := time.Parse("15:04", z.End)

			if !slotHour.Before(startTime) && slotHour.Before(endTime) {
				zones[z.Name] = append(zones[z.Name], model.Timeslot{
					Start: slot.StartTime,
					End:   slot.EndTime,
				})
				break
			}
		}
	}

	log.Printf("groupByZones summary zones: %d", len(zones))
	return zones
}

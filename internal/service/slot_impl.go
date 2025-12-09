package service

import (
	"fmt"
	"log"
	"time"

	rd "github.com/pan-asovsky/brandd-tg-bot/internal/cache/locker"
	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
	pg "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
)

type slotService struct {
	slotRepo   pg.SlotRepo
	slotLocker *rd.SlotLocker
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
		return nil, err
	}

	keys := make([]string, len(slots))
	for i, slot := range slots {
		keys[i] = s.slotLocker.FormatKey(slot.Date, fmt.Sprintf("%s-%s", slot.StartTime, slot.EndTime))
	}

	lockStatus, err := s.slotLocker.AreLocked(keys...)
	if err != nil {
		return nil, fmt.Errorf("[get_available_zones] failed to check locks: %w", err)
	}

	filtered := make([]model.Slot, 0, len(slots))
	for i, key := range keys {
		if !lockStatus[key] {
			filtered = append(filtered, slots[i])
		}
	}
	log.Printf("[get_available_zones] slots: %d, filtered: %d", len(slots), len(filtered))

	return s.groupByZones(filtered), nil
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

func (s *slotService) FindByDateAndTime(date, start string) (*model.Slot, error) {
	slot, err := s.slotRepo.FindByDateAndTime(date, start)
	if err != nil {
		return nil, fmt.Errorf("[find_by_date_time] error: %w", err)
	}
	return slot, nil
}

func (s *slotService) MarkUnavailable(date, start, end string) error {
	if err := s.slotRepo.MarkUnavailable(date, start, end); err != nil {
		return fmt.Errorf("[mark_unavailable] %w", err)
	}
	return nil
}

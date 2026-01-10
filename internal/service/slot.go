package service

import (
	"fmt"
	"log"
	"time"

	consts "github.com/pan-asovsky/brandd-tg-bot/internal/constants"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	irepo "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
	icache "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/service"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type slotService struct {
	slotRepo   irepo.SlotRepo
	slotLocker icache.SlotLocking
}

func (s *slotService) GetAvailableBookings() []entity.AvailableBooking {
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

	bookings := make([]entity.AvailableBooking, 0, 3)
	for _, day := range days {
		if !day.available {
			continue
		}
		bookings = append(bookings, entity.AvailableBooking{
			Date:  today.AddDate(0, 0, day.offset),
			Label: day.label,
		})
	}

	return bookings
}

func (s *slotService) GetAvailableZones(date string) (entity.Zone, error) {
	slots, err := s.slotRepo.GetAvailableSlots(date)
	if err != nil {
		return nil, err
	}

	keys := make([]string, len(slots))
	for x, slot := range slots {
		keys[x] = s.slotLocker.FormatKey(slot.Date, fmt.Sprintf("%s-%s", slot.StartTime, slot.EndTime))
	}

	lockStatus, err := s.slotLocker.AreLocked(keys...)
	if err != nil {
		return nil, utils.WrapError(err)
	}

	filtered := make([]entity.Slot, 0, len(slots))
	for x, key := range keys {
		if !lockStatus[key] {
			filtered = append(filtered, slots[x])
		}
	}
	log.Printf("[get_available_zones] slots: %d, filtered: %d", len(slots), len(filtered))

	return s.groupByZones(filtered), nil
}

func (s *slotService) groupByZones(slots []entity.Slot) entity.Zone {
	zones := make(entity.Zone)

	for _, slot := range slots {
		if !slot.IsAvailable {
			continue
		}

		slotTime, err := time.Parse("15:04", slot.StartTime)
		if err != nil {
			log.Printf("[group_by_zones] error parsing slot time %s: %v", slot.StartTime, err)
			continue
		}

		for _, z := range entity.ZonesDefinition {
			startTime, _ := time.Parse("15:04", z.Start)
			endTime, _ := time.Parse("15:04", z.End)

			if !slotTime.Before(startTime) && slotTime.Before(endTime) {
				zones[z.Name] = append(zones[z.Name], entity.Timeslot{
					Start: slot.StartTime,
					End:   slot.EndTime,
				})
				break
			}
		}
	}
	return zones
}

func (s *slotService) FindByDateAndTime(date, start string) (*entity.Slot, error) {
	return utils.WrapFunction(func() (*entity.Slot, error) {
		return s.slotRepo.FindByDateAndTime(date, start)
	})
}

func (s *slotService) MarkUnavailable(date, startTime string) error {
	return utils.WrapFunctionError(func() error {
		return s.slotRepo.MarkUnavailable(date, startTime)
	})
}

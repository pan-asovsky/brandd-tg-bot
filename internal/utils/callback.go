package utils

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"

	"github.com/pan-asovsky/brandd-tg-bot/internal/handler/types"
)

func GetSessionInfo(callback string) (*types.UserSessionInfo, error) {
	log.Printf("[parse_callback] raw callback: %s", callback)

	switch {
	case isDateCallback(callback):
		return parseDateCallback(callback)
	case isZoneCallback(callback):
		return parseZoneCallback(callback)
	case isTimeCallback(callback):
		return parseTimeCallback(callback)
	case isServiceCallback(callback):
		return parseServiceCallback(callback)
	case isRimCallback(callback):
		return parseRimCallback(callback)
	default:
		log.Printf("[parse_callback] unknown callback format: %s", callback)
		return &types.UserSessionInfo{}, nil
	}
}

func isDateCallback(cd string) bool {
	re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	return re.MatchString(cd)
}

func isZoneCallback(cd string) bool {
	re := regexp.MustCompile(`^\d{2}:\d{2}-\d{2}:\d{2}/\d{4}-\d{2}-\d{2}$`)
	return re.MatchString(cd)
}

func isTimeCallback(cd string) bool {
	re := regexp.MustCompile(`^\d{2}:\d{2}-\d{2}:\d{2}/\d{2}:\d{2}-\d{2}:\d{2}/\d{4}-\d{2}-\d{2}$`)
	return re.MatchString(cd)
}

func isServiceCallback(cd string) bool {
	re := regexp.MustCompile(`^[A-Z_]+(\+[A-Z_]+)*/\d{2}:\d{2}-\d{2}:\d{2}/\d{4}-\d{2}-\d{2}$`)
	return re.MatchString(cd)
}

func isRimCallback(cd string) bool {
	re := regexp.MustCompile(`^\d+/[A-Z_]+/\d{2}:\d{2}-\d{2}:\d{2}/\d{4}-\d{2}-\d{2}$`)
	return re.MatchString(cd)
}

func parseDateCallback(cd string) (*types.UserSessionInfo, error) {
	return &types.UserSessionInfo{Date: cd}, nil
}

func parseZoneCallback(cd string) (*types.UserSessionInfo, error) {
	parts := strings.Split(cd, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("[parse_zone] %w", errorInvalidPartsCount(cd, parts))
	}
	return &types.UserSessionInfo{Zone: parts[0], Date: parts[1]}, nil
}

func parseTimeCallback(cd string) (*types.UserSessionInfo, error) {
	parts := strings.Split(cd, "/")
	if len(parts) != 3 {
		return nil, fmt.Errorf("[parse_time] %w", errorInvalidPartsCount(cd, parts))
	}
	return &types.UserSessionInfo{Time: parts[0], Zone: parts[1], Date: parts[2]}, nil
}

func parseServiceCallback(cd string) (*types.UserSessionInfo, error) {
	parts := strings.Split(cd, "/")
	if len(parts) != 3 {
		return nil, fmt.Errorf("[parse_service] %w", errorInvalidPartsCount(cd, parts))
	}

	service := parts[0]
	if !isValidService(service) {
		return nil, errorInvalidService(service)
	}

	services := strings.Split(service, "+")
	sort.Strings(services)
	mappedService := mapServices(services)
	log.Printf("[parse_service] raw services: %v, mapped: %s", services, mappedService)

	return &types.UserSessionInfo{
		Service: mappedService,
		Time:    parts[1],
		Date:    parts[2],
	}, nil
}

func parseRimCallback(cd string) (*types.UserSessionInfo, error) {
	parts := strings.Split(cd, "/")
	if len(parts) != 4 {
		return nil, fmt.Errorf("[parse_rim] %w", errorInvalidPartsCount(cd, parts))
	}

	service := parts[1]
	if !isValidService(service) {
		return nil, errorInvalidService(service)
	}

	return &types.UserSessionInfo{
		Radius:  parts[0],
		Service: service,
		Time:    parts[2],
		Date:    parts[3],
	}, nil
}

func errorInvalidPartsCount(cd string, parts []string) error {
	return fmt.Errorf("[parse_callback] invalid parts count: %d, callback: %s", len(parts), cd)
}

func errorInvalidService(service string) error {
	return fmt.Errorf("[invalid_service] %s", service)
}

func mapServices(services []string) string {
	sort.Strings(services)
	servicesStr := strings.Join(services, "+")

	validCombinations := map[string]string{
		"TAKE_IT_OUT+BALANCING+TIRE_SERVICE": "COMPLEX",
		"BALANCING+TAKE_IT_OUT+TIRE_SERVICE": "COMPLEX",
		"TAKE_IT_OUT+TIRE_SERVICE+BALANCING": "COMPLEX",
		"TIRE_SERVICE+BALANCING+TAKE_IT_OUT": "COMPLEX",
		"BALANCING+TIRE_SERVICE+TAKE_IT_OUT": "COMPLEX",
		"TIRE_SERVICE+TAKE_IT_OUT+BALANCING": "COMPLEX",

		"TAKE_IT_OUT+BALANCING": "TAKE_AND_BALANCING",
		"BALANCING+TAKE_IT_OUT": "TAKE_AND_BALANCING",

		"TAKE_IT_OUT+TIRE_SERVICE": "TAKE_AND_TIRE",
		"TIRE_SERVICE+TAKE_IT_OUT": "TAKE_AND_TIRE",

		"BALANCING+TIRE_SERVICE": "TIRE_AND_BALANCING",
		"TIRE_SERVICE+BALANCING": "TIRE_AND_BALANCING",

		"TAKE_IT_OUT":  "TAKE_IT_OUT",
		"BALANCING":    "BALANCING",
		"TIRE_SERVICE": "TIRE_SERVICE",
		"COMPLEX":      "COMPLEX",
	}

	if mapped, ok := validCombinations[servicesStr]; ok {
		return mapped
	}

	log.Printf("[map_services] unknown combination: %s", servicesStr)
	return servicesStr
}

func isValidService(service string) bool {
	services := strings.Split(service, "+")

	validSingleServices := map[string]bool{
		"TAKE_IT_OUT":  true,
		"BALANCING":    true,
		"TIRE_SERVICE": true,
		"COMPLEX":      true,
	}

	for _, svc := range services {
		if svc == "" {
			log.Printf("[validate_service] empty service part")
			return false
		}

		re := regexp.MustCompile(`^[A-Z_]+$`)
		if !re.MatchString(svc) {
			log.Printf("[validate_service] invalid service syntax: %s", svc)
			return false
		}

		if !validSingleServices[svc] {
			log.Printf("[validate_service] unknown service: %s", svc)
			return false
		}
	}

	seen := make(map[string]bool)
	for _, svc := range services {
		if seen[svc] {
			log.Printf("[validate_service] duplicate service: %s", svc)
			return false
		}
		seen[svc] = true
	}

	if len(services) > 1 {
		sortedServices := make([]string, len(services))
		copy(sortedServices, services)
		sort.Strings(sortedServices)
		servicesStr := strings.Join(sortedServices, "+")

		validCombinations := []string{
			"TAKE_IT_OUT+BALANCING+TIRE_SERVICE",
			"TAKE_IT_OUT+BALANCING",
			"TAKE_IT_OUT+TIRE_SERVICE",
			"BALANCING+TIRE_SERVICE",
		}

		valid := false
		for _, comb := range validCombinations {
			if comb == servicesStr {
				valid = true
				break
			}
		}

		if !valid {
			log.Printf("[validate_service] invalid combination: %s", servicesStr)
			return false
		}
	}

	return len(services) > 0
}

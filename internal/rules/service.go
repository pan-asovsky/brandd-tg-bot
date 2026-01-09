package rules

import (
	"log"
	"sort"
	"strings"
)

const (
	Complex          = "COMPLEX"
	TakeItOut        = "TAKE_IT_OUT"
	TireService      = "TIRE_SERVICE"
	Balancing        = "BALANCING"
	TakeAndBalancing = "TAKE_AND_BALANCING"
	TakeAndTire      = "TAKE_AND_TIRE"
	TireAndBalancing = "TIRE_AND_BALANCING"
	PLUS             = "+"
)

type ServiceRules struct{}

func (r *ServiceRules) Apply(currentMap map[string]bool, clickedService string) map[string]bool {
	result := make(map[string]bool)
	for k, v := range currentMap {
		result[k] = v
	}

	if result[TakeItOut] && result[TireService] && result[Balancing] {
		return map[string]bool{Complex: true}
	}

	if clickedService == Complex {
		if result[Complex] {
			return map[string]bool{Complex: true}
		}
		return result
	}

	if result[Complex] {
		newMap := make(map[string]bool)
		newMap[clickedService] = true
		return newMap
	}

	return result
}

func (r *ServiceRules) MapServices(services []string) string {
	sort.Strings(services)
	servicesStr := strings.Join(services, "+")

	validCombinations := map[string]string{
		TakeItOut + PLUS + Balancing + PLUS + TireService: Complex,
		Balancing + PLUS + TakeItOut + PLUS + TireService: Complex,
		TakeItOut + PLUS + TireService + PLUS + Balancing: Complex,
		TireService + PLUS + Balancing + PLUS + TakeItOut: Complex,
		Balancing + PLUS + TireService + PLUS + TakeItOut: Complex,
		TireService + PLUS + TakeItOut + PLUS + Balancing: Complex,

		TakeItOut + PLUS + Balancing: TakeAndBalancing,
		Balancing + PLUS + TakeItOut: TakeAndBalancing,

		TakeItOut + PLUS + TireService: TakeAndTire,
		TireService + PLUS + TakeItOut: TakeAndTire,

		Balancing + PLUS + TireService: TireAndBalancing,
		TireService + PLUS + Balancing: TireAndBalancing,

		TakeAndBalancing + PLUS + TireService: Complex,
		TakeAndTire + PLUS + Balancing:        Complex,
		TireAndBalancing + PLUS + TakeItOut:   Complex,

		TakeItOut:   TakeItOut,
		Balancing:   Balancing,
		TireService: TireService,
		Complex:     Complex,
	}

	if mapped, ok := validCombinations[servicesStr]; ok {
		log.Printf("[map_services] in: %v, out: %s", services, mapped)
		return mapped
	}

	log.Printf("[map_services] unknown combination: %s", servicesStr)
	return servicesStr
}

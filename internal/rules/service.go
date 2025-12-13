package rules

import "log"

const (
	Complex     = "COMPLEX"
	TakeItOut   = "TAKE_IT_OUT"
	TireService = "TIRE_SERVICE"
	Balancing   = "BALANCING"
)

type ServiceRules struct{}

func (r *ServiceRules) Apply(currentMap map[string]bool, clickedService string) map[string]bool {
	log.Printf("[service_rules] current %#v, clicked %s", currentMap, clickedService)
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

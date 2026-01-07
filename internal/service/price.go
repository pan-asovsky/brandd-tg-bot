package service

import (
	"log"
	"strings"

	pg "github.com/pan-asovsky/brandd-tg-bot/internal/repository/postgres"
	"github.com/pan-asovsky/brandd-tg-bot/internal/utils"
)

type PriceService interface {
	Calculate(service, radius string) (int64, error)
}

type priceService struct {
	pgProvider *pg.Provider
}

func (p *priceService) Calculate(service, radius string) (int64, error) {
	services := splitServices(service)

	var totalPrice int64
	for _, svc := range services {
		price, err := p.pgProvider.Price().GetSetPrice(svc, radius)
		if err != nil {
			return 0, utils.WrapError(err)
		}
		totalPrice += price
	}

	return roundToFifty(totalPrice), nil
}

func splitServices(service string) []string {
	if strings.Contains(service, "AND") {
		switch service {
		case "TAKE_AND_BALANCING":
			return []string{"TAKE_IT_OUT", "BALANCING"}
		case "TAKE_AND_TIRE":
			return []string{"TAKE_IT_OUT", "TIRE_SERVICE"}
		case "TIRE_AND_BALANCING":
			return []string{"BALANCING", "TIRE_SERVICE"}
		default:
			log.Printf("unknown service: %s", service)
			return []string{}
		}
	}
	return []string{service}
}

func roundToFifty(x int64) int64 {
	return ((x + 24) / 50) * 50
}

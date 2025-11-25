package price

import "github.com/pan-asovsky/brandd-tg-bot/internal/model"

func calculate(wheelCount int64, price model.Price) int64 {
	if wheelCount == 4 {
		return price.PricePerSet
	}
	return price.PricePerWheel * wheelCount
}

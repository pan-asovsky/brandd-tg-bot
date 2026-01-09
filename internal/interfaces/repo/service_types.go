package interfaces

import "github.com/pan-asovsky/brandd-tg-bot/internal/model"

type ServiceRepo interface {
	GetServiceTypes() ([]model.ServiceType, error)
}

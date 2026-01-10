package interfaces

import "github.com/pan-asovsky/brandd-tg-bot/internal/entity"

type ServiceRepo interface {
	GetServiceTypes() ([]entity.ServiceType, error)
}

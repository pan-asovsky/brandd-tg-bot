package postgres

import (
	"database/sql"
	"fmt"

	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type ServiceRepo interface {
	GetServiceTypes() ([]model.ServiceType, error)
}

type serviceRepo struct {
	db *sql.DB
}

func NewServiceRepo(db *sql.DB) ServiceRepo {
	return &serviceRepo{db: db}
}

func (sr serviceRepo) GetServiceTypes() ([]model.ServiceType, error) {
	rows, err := sr.db.Query(GetAllServiceTypes)
	if err != nil {
		return nil, fmt.Errorf("[get_service_types] error: %w", err)
	}
	defer rows.Close()

	var types []model.ServiceType
	for rows.Next() {
		var svc model.ServiceType
		if err := rows.Scan(
			&svc.ID,
			&svc.ServiceCode,
			&svc.ServiceName,
			&svc.Description,
		); err != nil {
			return nil, fmt.Errorf("[get_service_types] rows scan error: %w", err)
		}
		types = append(types, svc)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[get_service_types] rows error: %w", err)
	}

	return types, nil
}

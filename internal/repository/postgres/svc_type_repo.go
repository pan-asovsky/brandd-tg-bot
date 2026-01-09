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

func (sr *serviceRepo) GetServiceTypes() ([]model.ServiceType, error) {
	rows, err := sr.db.Query(GetCompositeServiceTypes)
	if err != nil {
		return nil, fmt.Errorf("[get_service_types] query error: %w", err)
	}
	defer rows.Close()

	var types []model.ServiceType
	for rows.Next() {
		var svc model.ServiceType
		if err := rows.Scan(
			&svc.ID,
			&svc.ServiceCode,
			&svc.ServiceName,
			&svc.IsComposite,
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

package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type ServiceRepo interface {
	GetServiceTypes() ([]model.ServiceType, error)
	FindByCode(service string) (*model.ServiceType, error)
}

type serviceRepo struct {
	db *sql.DB
}

func (sr *serviceRepo) GetServiceTypes() ([]model.ServiceType, error) {
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

func (sr *serviceRepo) FindByCode(svc string) (*model.ServiceType, error) {
	var service model.ServiceType
	if err := sr.db.QueryRow(GetServiceTypeByCode, svc).Scan(
		&service.ID,
		&service.ServiceCode,
		&service.ServiceName,
		&service.Description,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("[find_by_code] service type not founded: %w", err)
		}
		return nil, fmt.Errorf("[find_by_code] failed: %w", err)
	}

	return &service, nil
}

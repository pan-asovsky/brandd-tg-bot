package postgres

import (
	"database/sql"
	"fmt"

	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	irepo "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
)

type serviceRepo struct {
	db *sql.DB
}

func NewServiceRepo(db *sql.DB) irepo.ServiceRepo {
	return &serviceRepo{db: db}
}

func (sr *serviceRepo) GetServiceTypes() ([]entity.ServiceType, error) {
	rows, err := sr.db.Query(GetCompositeServiceTypes)
	if err != nil {
		return nil, fmt.Errorf("[get_service_types] query error: %w", err)
	}
	defer rows.Close()

	var types []entity.ServiceType
	for rows.Next() {
		var svc entity.ServiceType
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

package postgres

import (
	"database/sql"
	"fmt"

	"github.com/pan-asovsky/brandd-tg-bot/internal/model"
)

type userRepo struct {
	db *sql.DB
}

func (u *userRepo) GetActiveAdmins() ([]model.User, error) {
	rows, err := u.db.Query(GetActiveAdmins)
	if err != nil {
		return nil, fmt.Errorf("[get_active_admins] query error: %w", err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.ID,
			&user.ChatID,
			&user.Name,
			&user.Role,
			&user.IsActive,
		); err != nil {
			return nil, fmt.Errorf("[get_active_admins] row scan error: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[get_active_admins] rows error: %w", err)
	}

	return users, nil
}

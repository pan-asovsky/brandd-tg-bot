package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	i "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) i.UserRepo {
	return &userRepo{db: db}
}

func (ur *userRepo) GetActiveAdmins() ([]entity.User, error) {
	rows, err := ur.db.Query(GetActiveAdmins)
	if err != nil {
		return nil, fmt.Errorf("[get_active_admins] query error: %w", err)
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
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

func (ur *userRepo) GetRole(chatID int64) (bool, string) {
	var userRole string
	err := ur.db.QueryRow(GetUserRole, chatID).Scan(&userRole)
	if err != nil {
		log.Printf("[get_user_role] user not exists: %d", chatID)
		return false, ""
	}

	return true, userRole
}

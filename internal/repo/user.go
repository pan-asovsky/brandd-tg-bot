package repo

import (
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pan-asovsky/brandd-tg-bot/internal/entity"
	irepo "github.com/pan-asovsky/brandd-tg-bot/internal/interfaces/repo"
)

type userRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepo(p *pgxpool.Pool) irepo.UserRepo {
	return &userRepo{pool: p}
}

func (ur *userRepo) GetActiveAdmins() ([]entity.User, error) {
	ctx, cancel := CtxWithTimeout(TwoSec)
	defer cancel()

	rows, err := ur.pool.Query(ctx, GetActiveAdmins)
	if err != nil {
		return nil, fmt.Errorf("[get_active_admins] query error: %w", err)
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		if err = rows.Scan(
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

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("[get_active_admins] rows error: %w", err)
	}

	return users, nil
}

func (ur *userRepo) GetRole(chatID int64) (bool, string) {
	ctx, cancel := CtxWithTimeout(TwoSec)
	defer cancel()

	var userRole string
	err := ur.pool.QueryRow(ctx, GetUserRole, chatID).Scan(&userRole)
	if err != nil {
		log.Printf("[get_user_role] user not exists: %d", chatID)
		return false, ""
	}

	return true, userRole
}

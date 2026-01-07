package model

type User struct {
	ID       int64  `json:"id"`
	ChatID   int64  `json:"chat_id"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	IsActive bool   `json:"is_active"`
}

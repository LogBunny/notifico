package models

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt  int       `json:"created_at"`
	UpdatedAt  int       `json:"updated_at"`
	UserName   string    `json:"user_name"`
	Name       string    `json:"name"`
	Email      string    `gorm:"unique; not null" json:"email"`
	TelegramId string    `json:"telegram_id"`
}

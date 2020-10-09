package domain

import "time"

// Recode struct
type Recode struct {
	ID        int64      `json:"user_id" gorm:"primary_key"`
	Message   string     `json:"message" validate:"required"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

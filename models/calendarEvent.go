package models

import (
	"time"

	"github.com/google/uuid"
)

// -------------------- CalendarEvent --------------------
type CalendarEvent struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	Title       string    `gorm:"not null"`
	StartTime   time.Time `gorm:"not null"`
	EndTime     *time.Time
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Todos []Todo
}

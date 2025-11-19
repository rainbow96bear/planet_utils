package models

import (
	"time"

	"github.com/google/uuid"
)

// -------------------- Profile --------------------
type Profile struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID       uuid.UUID `gorm:"type:uuid;unique;not null"`
	Nickname     string    `gorm:"size:50;not null"`
	Bio          string
	ProfileImage string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

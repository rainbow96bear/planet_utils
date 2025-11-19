package models

import (
	"time"

	"github.com/google/uuid"
)

// -------------------- Todo --------------------
type Todo struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	EventID   uuid.UUID `gorm:"type:uuid;not null"`
	Content   string    `gorm:"not null"`
	IsDone    bool      `gorm:"default:false"`
	DueTime   *time.Time
	Metadata  map[string]interface{} `gorm:"type:jsonb;default:'{}'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

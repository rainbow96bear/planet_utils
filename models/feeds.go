package models

import (
	"time"

	"github.com/google/uuid"
)

// -------------------- Feed --------------------
type Feeds struct {
	ID        uuid.UUID              `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID              `gorm:"type:uuid;not null"`
	Content   string                 `gorm:"not null"`
	Metadata  map[string]interface{} `gorm:"type:jsonb;default:'{}'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

package models

import (
	"time"

	"github.com/google/uuid"
)

// -------------------- ChatMessage --------------------
type ChatMessages struct {
	ID        uuid.UUID              `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	RoomID    uuid.UUID              `gorm:"type:uuid;not null"`
	SenderID  uuid.UUID              `gorm:"type:uuid;not null"`
	Content   string                 `gorm:"not null"`
	Metadata  map[string]interface{} `gorm:"type:jsonb;default:'{}'"`
	CreatedAt time.Time
}

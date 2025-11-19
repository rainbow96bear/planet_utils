package models

import (
	"time"

	"github.com/google/uuid"
)

// -------------------- Like --------------------
type Like struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID     uuid.UUID `gorm:"type:uuid;not null"`
	TargetType string    `gorm:"type:todo_feed_enum;not null"` // enum: 'todo','feed'
	TargetID   uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt  time.Time
}

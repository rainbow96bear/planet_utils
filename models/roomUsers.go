package models

import (
	"time"

	"github.com/google/uuid"
)

type RoomUsers struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	RoomID   uuid.UUID `gorm:"type:uuid;not null"`
	UserID   uuid.UUID `gorm:"type:uuid;not null"`
	JoinedAt time.Time
}

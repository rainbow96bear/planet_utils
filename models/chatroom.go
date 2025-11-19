package models

import (
	"time"

	"github.com/google/uuid"
)

// -------------------- ChatRoom --------------------
type ChatRoom struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Type      string    `gorm:"type:room_type_enum;not null"` // enum: '1:1','group'
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time

	Messages []ChatMessage
	Users    []RoomUser
}

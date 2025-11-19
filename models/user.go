package models

import (
	"time"

	"github.com/google/uuid"
)

// -------------------- User --------------------
type User struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email         string    `gorm:"unique;not null"`
	OAuthProvider string
	OAuthID       string
	CreatedAt     time.Time
	UpdatedAt     time.Time

	Profile        Profile
	CalendarEvents []CalendarEvent
	Todos          []Todo
	Feeds          []Feed
	Likes          []Like
	Follows        []Follow      `gorm:"foreignKey:FollowerID"`
	ChatMessages   []ChatMessage `gorm:"foreignKey:SenderID"`
}

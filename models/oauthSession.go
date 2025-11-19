package models

import (
	"time"

	"github.com/google/uuid"
)

type OAuthSession struct {
	ID            uuid.UUID              `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OAuthProvider string                 `gorm:"size:50;not null"`
	OAuthID       string                 `gorm:"size:255;not null"`
	Email         string                 `gorm:"size:255"`
	ProfileData   map[string]interface{} `gorm:"type:jsonb;default:'{}'"`
	CreatedAt     time.Time
	ExpiresAt     time.Time
}

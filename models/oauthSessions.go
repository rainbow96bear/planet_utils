package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type OauthSessions struct {
	ID            uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	OauthProvider string         `gorm:"column:oauth_provider"`
	OauthID       string         `gorm:"column:oauth_id"`
	Email         string         `gorm:"column:email"`
	ProfileData   datatypes.JSON `gorm:"column:profile_data"` // ★ 중요!
	CreatedAt     time.Time      `gorm:"column:created_at"`
	ExpiresAt     time.Time      `gorm:"column:expires_at"`
}

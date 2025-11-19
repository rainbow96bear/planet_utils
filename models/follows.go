package models

import (
	"time"

	"github.com/google/uuid"
)

type Follow struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FollowerID  uuid.UUID `gorm:"type:uuid;not null"`
	FollowingID uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt   time.Time
}

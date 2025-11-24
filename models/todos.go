package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// -------------------- Todo --------------------
type Todos struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID   uuid.UUID `gorm:"type:uuid;not null"`
	EventID  uuid.UUID `gorm:"type:uuid;not null"`
	Content  string    `gorm:"not null"`
	IsDone   bool      `gorm:"default:false"`
	DueTime  *time.Time
	Metadata datatypes.JSONMap `gorm:"type:jsonb;default:'{}'"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User  *Users          `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Event *CalendarEvents `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

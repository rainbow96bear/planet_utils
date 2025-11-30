package models

import (
	"time"

	"github.com/google/uuid"
)

// -------------------- CalendarEvent --------------------
type CalendarEvents struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	Title       string    `gorm:"size:50;not null"`
	Emoji       string    `gorm:"size:10;default:'📝'"`
	Description string    `gorm:"type:text"`
	StartAt     time.Time `gorm:"not null"`
	EndAt       time.Time `gorm:"not null"`
	Visibility  string    `gorm:"type:text;default:'public';check:visibility IN ('public','friends','private')"`
	ImageURL    string    `gorm:"type:text"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User  *Users   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Todos []*Todos `gorm:"foreignKey:EventID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

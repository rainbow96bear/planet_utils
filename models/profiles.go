package models

import (
	"time"

	"github.com/google/uuid"
)

type Profiles struct {
	ID             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID         uuid.UUID `gorm:"type:uuid;unique;not null"`
	Nickname       string    `gorm:"size:50;not null;unique"`
	Bio            string    `gorm:"type:text"`
	ProfileImage   string    `gorm:"type:text"`
	FollowerCount  int       `gorm:"not null;default:0"`
	FollowingCount int       `gorm:"not null;default:0"`
	Theme          string    `gorm:"type:varchar(10);not null;default:'light'"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	User *Users `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

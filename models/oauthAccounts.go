package models

import (
	"time"

	"github.com/google/uuid"
)

type OauthAccounts struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID        uuid.UUID `gorm:"type:uuid;not null"`
	OauthProvider string    `gorm:"size:50;not null"`
	OauthID       string    `gorm:"size:255;not null"`
	Email         string    `gorm:"type:varchar(255)"`

	CreatedAt time.Time `gorm:"autoCreateTime"`

	User *Users `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

package models

import (
	"time"

	"github.com/google/uuid"
)

type UserTokens struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null;index"`           // users 테이블 FK
	Token     string    `gorm:"type:text;not null"`                 // Refresh Token 값
	TokenType string    `gorm:"type:varchar(20);default:'refresh'"` // 'refresh' / 'session'
	IPAddress string    `gorm:"size:50"`                            // Optional: 로그인 시 IP
	UserAgent string    `gorm:"type:text"`                          // Optional: 브라우저/앱 정보
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time
	Revoked   bool `gorm:"default:false"` // 토큰 회수 여부
}

package model

import (
	"time"
)

// User : 회원 정보 구조체
type User struct {
	ID             uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserUUID       string     `gorm:"uniqueIndex;size:36" json:"user_uuid"`
	OAuthPlatform  string     `gorm:"size:50" json:"oauth_platform"`
	OAuthID        string     `gorm:"size:100" json:"oauth_id"`
	Email          string     `gorm:"size:100" json:"email,omitempty"`
	Nickname       string     `gorm:"size:50" json:"nickname"`
	Bio            string     `gorm:"type:text" json:"bio,omitempty"`
	ProfileImage   string     `gorm:"size:255" json:"profile_image"`
	Role           string     `gorm:"size:20" json:"role"`
	Theme          string     `gorm:"size:20" json:"theme"`
	LastLogin      *time.Time `json:"last_login,omitempty"`
	FollowerCount  uint       `gorm:"default:0" json:"follower_count"`
	FollowingCount uint       `gorm:"default:0" json:"following_count"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

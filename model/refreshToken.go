package model

import "time"

type RefreshToken struct {
	ID        uint64    `json:"id" db:"id"`                 // AUTO_INCREMENT 기본키
	UserUUID  string    `json:"user_uuid" db:"user_uuid"`   // users.user_uuid 참조
	Token     string    `json:"token" db:"token"`           // 토큰 문자열
	Expiry    uint64    `json:"expiry" db:"expiry"`         // 만료 시간 (UNIX 타임스탬프)
	CreatedAt time.Time `json:"created_at" db:"created_at"` // 생성일
}

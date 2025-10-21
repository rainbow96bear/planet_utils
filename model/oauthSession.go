package model

import "time"

// OAuthSession : OAuth 인증 세션 정보 구조체
type OAuthSession struct {
	SessionID     string    `json:"session_id" db:"session_id"`         // 랜덤 토큰 (64자)
	OAuthPlatform string    `json:"oauth_platform" db:"oauth_platform"` // OAuth 제공자 (google, kakao, github 등)
	OAuthID       string    `json:"oauth_id" db:"oauth_id"`             // OAuth 제공자 고유 ID
	ExpiresAt     time.Time `json:"expires_at" db:"expires_at"`         // 만료 시간
	CreatedAt     time.Time `json:"created_at" db:"created_at"`         // 세션 생성 시간
}

package model

import "time"

// User : 회원 정보 구조체
type User struct {
	ID            uint64    `json:"id" db:"id"`                           // AUTO_INCREMENT 기본키
	UserUuid      string    `json:"user_uuid" db:"user_uuid"`             // 외부 노출용 UUID
	OAuthPlatform string    `json:"oauth_platform" db:"oauth_platform"`   // OAuth 제공자 (google, kakao 등)
	OAuthID       string    `json:"oauth_id" db:"oauth_id"`               // OAuth 제공자에서 받은 고유 ID
	Email         string    `json:"email,omitempty" db:"email"`           // 이메일 (NULL 가능)
	Nickname      string    `json:"nickname" db:"nickname"`               // 닉네임
	Bio           string    `json:"bio,omitempty" db:"bio"`               // 자기소개
	ProfileImage  string    `json:"profile_image" db:"profile_image"`     // 프로필 이미지 URL
	Role          string    `json:"role" db:"role"`                       // 권한 (user/admin)
	CreatedAt     time.Time `json:"created_at" db:"created_at"`           // 생성일
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`           // 수정일
	LastLogin     time.Time `json:"last_login,omitempty" db:"last_login"` // 마지막 로그인 (NULL 가능)
}

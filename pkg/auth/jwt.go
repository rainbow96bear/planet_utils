package auth

import (
	"crypto/rand"     // 🚨 Refresh Token 생성을 위해 추가
	"encoding/base64" // 🚨 Refresh Token 문자열 인코딩을 위해 추가
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// ----------------------------------------------------------------------
// 1. CustomClaims 정의 (JWT에 담을 정보)
// ----------------------------------------------------------------------

// CustomClaims는 Access Token에 담을 사용자 정의 클레임입니다.
// UserID와 표준 클레임(RegisteredClaims)을 포함합니다.
type CustomClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

type SignupClaims struct {
	Provider   string `json:"provider"`
	ProviderID string `json:"provider_id"`
	jwt.RegisteredClaims
}

// ----------------------------------------------------------------------
// 2. JWT 생성 함수 (Signing)
// ----------------------------------------------------------------------

// GenerateAccessToken은 주어진 사용자 ID, 비밀 키, 만료 시간을 기반으로 JWT를 생성합니다.
func GenerateAccessToken(userID uuid.UUID, secretKey string, expiry time.Time) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiry),     // 토큰 만료 시간 (exp)
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 토큰 발급 시간 (iat)
			NotBefore: jwt.NewNumericDate(time.Now()), // 토큰 활성화 시간 (nbf)
			Subject:   userID.String(),
		},
	}

	// 서명 알고리즘 HS256 사용
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 비밀 키를 사용하여 토큰에 서명합니다.
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign access token: %w", err)
	}

	return signedToken, nil
}

// ----------------------------------------------------------------------
// 3. JWT 검증 및 클레임 추출 함수 (Verification & Parsing)
// ----------------------------------------------------------------------

// ParseAndVerifyCustomClaims는 토큰 문자열을 파싱하고 서명을 검증한 후,
// 유효하면 CustomClaims 구조체를 반환합니다.
func ParseAndVerifyCustomClaims(tokenStr, secretKey string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		// 토큰의 서명 알고리즘이 HMAC인지 확인
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		// 서명 검증을 위해 비밀 키를 제공합니다.
		return []byte(secretKey), nil
	})

	if err != nil {
		// 파싱 중 발생하는 오류(예: 서명 불일치, 형식 오류 등) 처리
		return nil, fmt.Errorf("token parsing failed: %w", err)
	}

	// 토큰의 유효성 (만료 시간 등)을 최종 확인합니다.
	if !token.Valid {
		return nil, fmt.Errorf("token is invalid or expired")
	}

	// 성공적으로 파싱된 claims 포인터를 반환합니다.
	return claims, nil
}

// ----------------------------------------------------------------------
// 4. Refresh Token 값 생성 함수
// ----------------------------------------------------------------------

// RefreshTokenLength는 생성할 Refresh Token의 바이트 길이입니다.
const RefreshTokenLength = 32

// GenerateRefreshToken creates a cryptographically secure, random string for the refresh token value.
// Refresh Token은 DB에 저장될 값으로, JWT가 아닌 무작위 문자열을 사용합니다.
func GenerateRefreshToken() (string, error) {
	// 32 바이트의 무작위 데이터를 담을 버퍼 생성
	tokenBytes := make([]byte, RefreshTokenLength)

	// crypto/rand를 사용하여 안전한 무작위 데이터로 채웁니다.
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes for refresh token: %w", err)
	}

	// URL-safe Base64 인코딩을 사용하여 문자열로 변환합니다.
	refreshToken := base64.URLEncoding.EncodeToString(tokenBytes)

	return refreshToken, nil
}

func GenerateSignupToken(provider, providerID, secret string, expiry time.Time) (string, error) {

	claims := SignupClaims{
		Provider:   provider,
		ProviderID: providerID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiry),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signed, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign signup token: %w", err)
	}

	return signed, nil
}

func ParseAndVerifySignupClaims(tokenStr, secret string) (*SignupClaims, error) {
	claims := &SignupClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		// HS256 검증
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("signup token parsing failed: %w", err)
	}

	// 만료 등 최종 검증
	if !token.Valid {
		return nil, fmt.Errorf("signup token is invalid or expired")
	}

	return claims, nil
}

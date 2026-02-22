package planetjwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	secret []byte
	parser *jwt.Parser
}

func NewManager(secret string) (*Manager, error) {
	if secret == "" {
		return nil, errors.New("JWT secret required")
	}
	return &Manager{
		secret: []byte(secret),
		parser: jwt.NewParser(
			jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
			jwt.WithLeeway(5*time.Second),
		),
	}, nil
}

func (m *Manager) Generate(userUUID string, tokenType TokenType, expiry time.Duration) (string, error) {
	claims := Claims{
		UserUUID:  userUUID,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(m.secret)
}

func (m *Manager) Verify(tokenStr string, expected TokenType) (*Claims, error) {
	claims := &Claims{}
	_, err := m.parser.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return m.secret, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token expired")
		}
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	if claims.TokenType != expected {
		return nil, errors.New("invalid token type")
	}
	return claims, nil
}

// 편의 메서드 추가
func (m *Manager) VerifySignupToken(tokenStr string) (*Claims, error) {
	return m.Verify(tokenStr, TokenTypeSignup)
}

func (m *Manager) VerifyAccessToken(tokenStr string) (*Claims, error) {
	return m.Verify(tokenStr, TokenTypeAccess)
}

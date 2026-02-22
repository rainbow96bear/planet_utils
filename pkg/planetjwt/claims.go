package planetjwt

import "github.com/golang-jwt/jwt/v5"

type TokenType string

const (
	TokenTypeSignup TokenType = "signup"
	TokenTypeAccess TokenType = "access"
)

type Claims struct {
	UserUUID  string    `json:"user_id"`
	TokenType TokenType `json:"token_type"`
	jwt.RegisteredClaims
}

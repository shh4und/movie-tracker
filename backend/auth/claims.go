package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims includes standard claims and any additional claims
type CustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

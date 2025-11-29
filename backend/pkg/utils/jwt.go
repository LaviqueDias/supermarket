// ============================================
// pkg/utils/jwt.go
// ============================================
package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/LaviqueDias/supermarket/pkg/config"
)

type Claims struct {
	ClientID int    `json:"client_id"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(clientID int, email string) (string, error) {
	cfg := config.Get()
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		ClientID: clientID,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

func ValidateToken(tokenString string) (*Claims, error) {
	cfg := config.Get()
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token inválido")
	}

	return claims, nil
}
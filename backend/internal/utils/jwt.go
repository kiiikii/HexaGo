package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("super-secret-hexago-key")

func GenerateToken(userID string) (string, error) {
	expiredTime := time.Now().Add(24 * time.Hour)

	claims := jwt.RegisteredClaims{
		Subject:   userID,
		ExpiresAt: jwt.NewNumericDate(expiredTime),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokens, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokens, nil
}

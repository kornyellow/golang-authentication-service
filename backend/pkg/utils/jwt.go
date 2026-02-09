package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint, username string) (string, error) {
	// 1. Create Claims
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Expire in 24 hours
	}

	// 2. Create Token Object
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 3. Sign Token with Secret
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}

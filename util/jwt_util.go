package util

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken() (string, error) {
	mapClaims := jwt.MapClaims{
		"email":      os.Getenv("USER_EMAIL"),
		"username":   os.Getenv("USER_USERNAME"),
		"created_at": time.Now(),
	}
	jtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	token, err := jtoken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}

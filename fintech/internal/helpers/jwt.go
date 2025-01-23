package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID uint) (*string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}

	token_jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := token_jwt.SignedString([]byte("TokenPassword"))
	if err != nil {
		return nil, err
	}

	return &token, nil
}

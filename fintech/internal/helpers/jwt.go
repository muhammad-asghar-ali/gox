package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID uint) (string, error) {
	token_c := jwt.MapClaims{
		"user_id": userID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}

	token_jwt := jwt.NewWithClaims(jwt.GetSigningMethod("256"), token_c)
	token, err := token_jwt.SignedString([]byte("TokenPassword"))
	if err != nil {
		return "", err
	}

	return token, nil
}

package common

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/muhammad-asghar-ali/gox/sysdesigns/ticketmaster/config"
)

func AccessToken(user_id uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user_id,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	return GenerateToken(claims)
}

func RefreshToken(user_id uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user_id,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	return GenerateToken(claims)
}

func GenerateToken(claims jwt.MapClaims) (string, error) {
	new := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := new.SignedString([]byte(config.GetConfig().GetJwtSecret()))
	if err != nil {
		return "", err
	}

	return token, nil
}

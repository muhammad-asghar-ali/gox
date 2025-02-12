package shared

import (
	"authis/internal/config"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func AccessToken(user_id uint) (*string, error) {
	claims := jwt.MapClaims{
		"user_id": user_id,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	return GenerateToken(claims)
}

func RefreshToken(user_id uint) (*string, error) {
	claims := jwt.MapClaims{
		"user_id": user_id,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	return GenerateToken(claims)
}

func GenerateToken(claims jwt.MapClaims) (*string, error) {
	new := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	config := config.GetConfig()

	token, err := new.SignedString([]byte(config.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func HashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash)
}

func ComparePassword(hash, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return err
	}

	return nil
}

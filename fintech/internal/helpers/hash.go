package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// hash
func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleError(err)

	return string(hashed)
}

func ComparePassword(pass, hash string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass)); err != nil {
		return err
	}

	return nil
}

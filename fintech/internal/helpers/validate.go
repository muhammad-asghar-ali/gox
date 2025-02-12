package helpers

import (
	"errors"
	"regexp"
)

type (
	Validate struct {
		Value string
		Valid string
	}
)

func Validation(vals []Validate) error {
	username := regexp.MustCompile(`^[A-Za-z0-9]{5,}$`)
	email := regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)

	for _, v := range vals {
		switch v.Valid {
		case "username":
			if !username.MatchString(v.Value) {
				return errors.New("invalid username: must be at least 5 alphanumeric characters")
			}
		case "email":
			if !email.MatchString(v.Value) {
				return errors.New("invalid email address")
			}
		case "password":
			if len(v.Value) < 5 {
				return errors.New("invalid password: must be at least 5 characters")
			}
		default:
			return errors.New("unknown validation type")
		}
	}

	return nil
}

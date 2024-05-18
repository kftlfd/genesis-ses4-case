package utils

import (
	"fmt"
	"net/mail"
)

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)

	if err != nil {
		return fmt.Errorf("invalid email: %s", email)
	}

	return nil
}

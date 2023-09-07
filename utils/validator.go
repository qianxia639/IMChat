package utils

import (
	"fmt"
	"net/mail"
)

func ValidatorEmail(value string) error {
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email address")
	}
	return nil
}

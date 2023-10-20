package utils

import (
	"net/mail"
)

func ValidatorEmail(value string) error {
	if _, err := mail.ParseAddress(value); err != nil {
		return err
	}
	return nil
}

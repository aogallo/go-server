package utils

import "golang.org/x/crypto/bcrypt"

func HasPassword(password string) (string, error) {
	hassPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	return string(hassPassword), err
}

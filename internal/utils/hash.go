package utils

import "golang.org/x/crypto/bcrypt"

func HasPassword(password string) (string, error) {
	hassPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	return string(hassPassword), err
}

func ComparePasswor(hashedPassword []byte, password string) bool {
	result := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))

	if result != nil {
		return false
	}

	return true
}

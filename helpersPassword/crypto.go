package helpersPassword

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func Decrypt(encrypted, entered string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(entered)); err != nil {
		return err
	}
	return nil
}

package util

import "golang.org/x/crypto/bcrypt"

func HashPassword(rawpassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawpassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(rawpassword, hashedPassword string) error {
	valid := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawpassword))
	return valid
}

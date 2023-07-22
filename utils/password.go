package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(pass), nil
}

func PasswordCheck(password string, hashPassword string) (err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)); err != nil {
		return
	}
	return
}

package utils_bcrypt

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func EncryptString(input string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input), bcrypt.DefaultCost)

	if err != nil {
		log.Println("[utils][bcrypt][EncryptString] error encrypting string", err.Error())
		return "", err
	}

	return string(hashedPassword), nil
}

func CompareHashAndPassword(inputData string, comparedData string) error {
	errPassword := bcrypt.CompareHashAndPassword([]byte(inputData), []byte(comparedData))

	if errPassword != nil {
		log.Println("[utils][bcrypt][EncryptString] error encrypting string", errPassword.Error())
		return errPassword
	}

	return nil
}
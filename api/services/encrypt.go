package services

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	encryptText, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	return string(encryptText), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	fmt.Println(err)

	return err == nil
}

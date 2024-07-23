package services

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func DecodeToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok && token.Valid {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		secret := []byte(os.Getenv("SECRET_KEY"))

		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, err
	}
}

func GenerateToken(data map[string]string) (string, error) {
	dataToken := jwt.MapClaims{}
	secretKey := []byte(os.Getenv("SECRET_KEY"))

	for k, v := range data {
		dataToken[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, dataToken)

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

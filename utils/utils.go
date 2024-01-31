package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/junanda/golang-aa/models"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CompatreHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ParseToken(tokenString string) (claim *models.Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})

	if err != nil {
		return nil, err
	}

	claim, ok := token.Claims.(*models.Claims)
	if !ok {
		return nil, err
	}

	return claim, nil
}

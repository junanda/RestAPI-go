package utils

import (
	"log"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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

	claim, ok := token.Claims.(*models.Claims)
	if !ok {
		log.Println("claims extract: ", err.Error())
		return claim, err
	}

	if err != nil {
		log.Println("parsing with claim: ", err.Error())
		return claim, err
	}

	return claim, nil
}

func GenerateUUID() string {
	idString := uuid.New()
	return regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(idString.String(), "")
}

func GetTokenString(ctx *gin.Context) string {
	header_auth := ctx.Request.Header.Get("Authorization")
	token := ""
	if len(strings.Split(header_auth, " ")) == 2 {
		token = strings.Split(header_auth, " ")[1]
	}
	return token
}

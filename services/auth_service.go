package services

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/junanda/golang-aa/models"
	"github.com/junanda/golang-aa/repository"
	"github.com/junanda/golang-aa/utils"
)

type AuthService interface {
	ValidationJwt(ctx *gin.Context, token string) (*models.Claims, error)
}

type authService struct {
	authRepo repository.AuthRepository
}

func NewAuthService(ar repository.AuthRepository) AuthService {
	return &authService{
		authRepo: ar,
	}
}

func (a *authService) ValidationJwt(ctx *gin.Context, token string) (*models.Claims, error) {
	var (
		claims *models.Claims
	)
	claims, err := utils.ParseToken(token)
	if err != nil {
		v, _ := err.(*jwt.ValidationError)
		switch v.Errors {
		case jwt.ValidationErrorSignatureInvalid:
			return nil, errors.New("Unauthorized")
		case jwt.ValidationErrorExpired:
			return nil, errors.New("Unauthorized, Token expired")
		default:
			return nil, errors.New("Unauthorized")
		}
	}

	authData, err := a.authRepo.GetData(ctx, token)
	if err != nil {
		log.Println("error data redis : ", err.Error())
		return nil, errors.New("Unauthorized")
	}

	if authData == (models.AuthData{}) {
		log.Println("user not found")
		return nil, errors.New("Unauthorized")
	}

	return claims, nil
}

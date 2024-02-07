package services

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/junanda/golang-aa/utils"
)

type BerandaService interface {
	Beranda(ctx *gin.Context, token string) (string, error)
	Premium(ctx *gin.Context, token string) (string, error)
}

type BerandaServiceImpl struct{}

func Init() BerandaService {
	return &BerandaServiceImpl{}
}

func (bs *BerandaServiceImpl) Beranda(ctx *gin.Context, token string) (string, error) {
	claims, err := utils.ParseToken(token)
	if err != nil {
		return "", errors.New("user unauthorized")
	}

	if claims.Role != "user" && claims.Role != "admin" {
		return "", errors.New("user unauthorized")
	}

	return claims.Role, nil
}

func (bs *BerandaServiceImpl) Premium(ctx *gin.Context, token string) (string, error) {
	claims, err := utils.ParseToken(token)
	if err != nil {
		return "", errors.New("user unauthorized")
	}

	if claims.Role != "admin" {
		return "", errors.New("user unauthorized")
	}

	return claims.Role, nil
}

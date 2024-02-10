package services

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type BerandaService interface {
	Beranda(ctx *gin.Context) (string, error)
	Premium(ctx *gin.Context) (string, error)
}

type BerandaServiceImpl struct{}

func Init() BerandaService {
	return &BerandaServiceImpl{}
}

func (bs *BerandaServiceImpl) Beranda(ctx *gin.Context) (string, error) {
	role := ctx.GetString("role")

	if role != "user" && role != "admin" {
		return "", errors.New("user unauthorized")
	}

	return role, nil
}

func (bs *BerandaServiceImpl) Premium(ctx *gin.Context) (string, error) {
	role := ctx.GetString("role")

	if role != "admin" {
		return "", errors.New("user unauthorized")
	}

	return role, nil
}

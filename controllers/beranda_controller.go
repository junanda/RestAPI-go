package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/junanda/golang-aa/middleware"
	"github.com/junanda/golang-aa/services"
)

type BerandaController interface {
	Handler(r *gin.Engine)
}

type BerandaControllerImpl struct {
	md       middleware.HeaderMiddleWare
	berandaS services.BerandaService
}

func Init(sb services.BerandaService, amd middleware.HeaderMiddleWare) BerandaController {
	return &BerandaControllerImpl{
		md:       amd,
		berandaS: sb,
	}
}

func (b *BerandaControllerImpl) Handler(r *gin.Engine) {
	berandaRoute := r.Group("/beranda", b.md.IsAuthorized)
	{
		berandaRoute.GET("/", func(ctx *gin.Context) {
			// cookie, err := ctx.Cookie("token")
			// if err != nil {
			// 	ctx.JSON(http.StatusUnauthorized, gin.H{
			// 		"message": "unautorized",
			// 	})

			// 	return
			// }

			role, err := b.berandaS.Beranda(ctx)
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": err.Error(),
				})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"message": "Home page",
				"role":    role,
			})

		})

		berandaRoute.GET("/premium", func(ctx *gin.Context) {
			// cookie, err := ctx.Cookie("token")
			// if err != nil {
			// 	ctx.JSON(http.StatusUnauthorized, gin.H{
			// 		"message": "unauthorized",
			// 	})
			// 	return
			// }

			role, err := b.berandaS.Premium(ctx)
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": err.Error(),
				})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"message": "Premium page",
				"role":    role,
			})
		})
	}
}

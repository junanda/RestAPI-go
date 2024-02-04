package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/junanda/golang-aa/models"
	"github.com/junanda/golang-aa/services"
)

type AuthController interface {
	Handler(r *gin.Engine)
}

type AuthControllerImpl struct {
	userService services.UserService
}

func InitAuthController(us services.UserService) AuthController {
	return &AuthControllerImpl{
		userService: us,
	}
}

func (a *AuthControllerImpl) Handler(r *gin.Engine) {
	authRoute := r.Group("/auth")
	{
		authRoute.POST("/login", func(ctx *gin.Context) {
			var (
				user models.User
			)

			if err := ctx.ShouldBindJSON(&user); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			err := a.userService.LoginUser(ctx, user)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"success": "user logged",
			})
		})

		authRoute.POST("/signup", func(ctx *gin.Context) {
			var (
				user models.User
			)

			if err := ctx.ShouldBindJSON(&user); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			err := a.userService.SignUp(ctx, user)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
			}
		})

		authRoute.GET("/logout", func(ctx *gin.Context) {
			log.Println("Logout route")
		})
	}
}

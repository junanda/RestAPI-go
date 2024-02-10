package controllers

import (
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

			token, err := a.userService.LoginUser(ctx, user)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"success": "user logged",
				"token":   token,
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

			ctx.JSON(http.StatusOK, gin.H{
				"message": "user registered",
			})
		})

		authRoute.GET("/logout", func(ctx *gin.Context) {
			err := a.userService.LogOut(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"failed": err.Error()})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{"success": "user logged out"})
		})
	}
}

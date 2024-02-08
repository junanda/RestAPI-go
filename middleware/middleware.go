package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/junanda/golang-aa/utils"
)

func IsAuthorized() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookies, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})

			ctx.Abort()
			return
		}

		claims, err := utils.ParseToken(cookies)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "user unauthorized",
			})

			ctx.Abort()
			return
		}

		ctx.Set("role", claims.Role)
		ctx.Next()
	}
}

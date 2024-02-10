package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/junanda/golang-aa/utils"
)

func IsAuthorized() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// cookies, err := ctx.Cookie("token")
		// if err != nil {
		// 	ctx.JSON(http.StatusUnauthorized, gin.H{
		// 		"message": "unauthorized",
		// 	})

		// 	ctx.Abort()
		// 	return
		// }

		header_auth := ctx.Request.Header.Get("Authorization")
		token := ""
		if len(strings.Split(header_auth, " ")) == 2 {
			token = strings.Split(header_auth, " ")[1]
		}

		claims, err := utils.ParseToken(token)
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

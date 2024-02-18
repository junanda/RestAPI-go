package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/junanda/golang-aa/repository"
	"github.com/junanda/golang-aa/utils"
)

type HeaderMiddleWare interface {
	IsAuthorized(ctx *gin.Context)
}

type headerMiddleware struct {
	authRepo repository.AuthRepository
}

func NewHeaderMiddleware(authRepo repository.AuthRepository) HeaderMiddleWare {
	return &headerMiddleware{
		authRepo: authRepo,
	}
}

func (h *headerMiddleware) IsAuthorized(ctx *gin.Context) {
	header_auth := ctx.Request.Header.Get("Authorization")
	token := ""
	if len(strings.Split(header_auth, " ")) == 2 {
		token = strings.Split(header_auth, " ")[1]
	}

	claim, err := utils.ParseToken(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "user unauthorized",
		})
		ctx.Abort()
		return
	}

	ctx.Set("role", claim.Role)
	ctx.Next()
}

func isAuthorized() gin.HandlerFunc {
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

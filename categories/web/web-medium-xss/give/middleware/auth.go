package middleware

import (
	"example/v3/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		token, err := ctx.Cookie("Authorization")
		if err != nil || token == "" {
			ctx.Set("isAuthenticated", false)
		} else {

			claims, err := utils.VerifyToken(token)
			if err != nil {
				ctx.Set("isAuthenticated", false)
			} else {
				ctx.Set("isAuthenticated", true)
				ctx.Set("claims", claims)
			}
		}
		ctx.Next()
	}
}
